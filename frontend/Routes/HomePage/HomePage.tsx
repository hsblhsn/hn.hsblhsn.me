import { DocumentNode, gql, useQuery } from '@apollo/client'
import { useLocation, useSearchParams } from 'react-router-dom'
import { Container, PaddedBlock } from '../../Components/Layout'
import { Block } from 'baseui/block'
import { NavBar } from '../../Components/NavBar'
import {
  PaginatedItemCardList,
  ITEM_CARD_LIST_STORY_FIELDS,
  ITEM_CARD_LIST_JOB_FIELDS,
} from '../../Components/ItemCardList'
import { ErrorScreen } from './ErrorScreen'
import { LoadingScreen } from './LoadingScreen'
import { ConnectionT, Job, Story } from '../../Types'
import { Fragment } from 'react'
import { Head } from './Head'

const PAGE_INFO_FIELDS = gql`
  fragment PageInfoFields on PageInfo {
    hasNextPage
    endCursor
    pageCursor
  }
`

const GET_TOP_STORIES = gql`
  ${ITEM_CARD_LIST_STORY_FIELDS}
  ${PAGE_INFO_FIELDS}
  query GetTopStories($after: Cursor) {
    items: topStories(after: $after, first: 9) {
      pageInfo {
        ...PageInfoFields
      }
      ...ItemCardListStoryFields
    }
  }
`

const GET_NEW_STORIES = gql`
  ${ITEM_CARD_LIST_STORY_FIELDS}
  ${PAGE_INFO_FIELDS}
  query GetNewStories($after: Cursor) {
    items: newStories(after: $after, first: 9) {
      pageInfo {
        ...PageInfoFields
      }
      ...ItemCardListStoryFields
    }
  }
`

const GET_ASK_STORIES = gql`
  ${ITEM_CARD_LIST_STORY_FIELDS}
  ${PAGE_INFO_FIELDS}
  query GetAskStories($after: Cursor) {
    items: askStories(after: $after, first: 9) {
      pageInfo {
        ...PageInfoFields
      }
      ...ItemCardListStoryFields
    }
  }
`

const GET_SHOW_STORIES = gql`
  ${ITEM_CARD_LIST_STORY_FIELDS}
  ${PAGE_INFO_FIELDS}
  query GetShowStories($after: Cursor) {
    items: showStories(after: $after, first: 9) {
      pageInfo {
        ...PageInfoFields
      }
      ...ItemCardListStoryFields
    }
  }
`

const GET_JOBS = gql`
  ${ITEM_CARD_LIST_JOB_FIELDS}
  ${PAGE_INFO_FIELDS}
  query GetJobStories($after: Cursor) {
    items: jobStories(after: $after) {
      pageInfo {
        ...PageInfoFields
      }
      ...ItemCardListJobFields
    }
  }
`

interface HomePageQueryData {
  items: ConnectionT<Story | Job>
}

interface HomePageQueryVars {
  after: string | null
}

const HomePage: React.FC = () => {
  const location = useLocation()
  let graphqlQuery = GET_TOP_STORIES
  switch (getBasePath(location.pathname)) {
    case '/':
      graphqlQuery = GET_TOP_STORIES
      break
    case '/newest':
      graphqlQuery = GET_NEW_STORIES
      break
    case '/ask':
      graphqlQuery = GET_ASK_STORIES
      break
    case '/show':
      graphqlQuery = GET_SHOW_STORIES
      break
    case '/jobs':
      graphqlQuery = GET_JOBS
      break
  }
  const head = <Head path={getBasePath(location.pathname)} />
  return (
    <Fragment>
      {head}
      <Container
        top={<NavBar />}
        left={<Block />}
        center={<InfiniteScroll query={graphqlQuery} />}
        right={<Block />}
      />
    </Fragment>
  )
}

const InfiniteScroll = ({ query }: { query: DocumentNode }) => {
  const [searchParams, setSearchParams] = useSearchParams()
  const { loading, error, data, fetchMore } = useQuery<
    HomePageQueryData,
    HomePageQueryVars
  >(query, {
    variables: {
      after: searchParams.get('after') || '0',
    },
    notifyOnNetworkStatusChange: true,
  })
  const loadNext = () => {
    fetchMore({
      variables: {
        after: data?.items.pageInfo.endCursor,
      },
    })
      .then(({ data }) => {
        if (data?.items.pageInfo?.pageCursor) {
          searchParams.set('after', data.items.pageInfo.pageCursor)
          setSearchParams(searchParams)
        }
      })
      .catch((e) => {
        console.error('Apollo error', e)
      })
  }
  let children: React.ReactNode = <Fragment />
  if (!data && loading) {
    children = <LoadingScreen />
  } else if (!data && error) {
    children = <ErrorScreen error={error} />
  } else if (data) {
    children = (
      <PaginatedItemCardList
        nextPageUrl={`?after=${data.items.pageInfo.pageCursor}`}
        loadNext={loadNext}
        loading={loading}
        items={data.items}
      />
    )
  }
  return <PaddedBlock>{children}</PaddedBlock>
}

function getBasePath(path: string): string {
  if (path) {
    return path.substring(path.lastIndexOf('/'), path.length)
  }
  return '/'
}

export { HomePage }
