import { Block } from 'baseui/block'
import { Container, PaddedBlock } from '../../Components/Layout'
import { useSearchParams } from 'react-router-dom'
import { LoadingScreen } from './LoadingScreen'
import { ErrorScreen } from './ErrorScreen'
import { FormEvent, Fragment, useCallback, useEffect, useState } from 'react'
import { BaseInput } from 'baseui/input'
import {
  ITEM_CARD_LIST_NODE_FIELDS,
  PaginatedItemCardList,
} from '../../Components/ItemCardList'
import { gql, useQuery } from '@apollo/client'
import { ConnectionT, Story } from '../../Types'
import { useStyletron } from 'baseui'
import { FormControl } from 'baseui/form-control'
import AwesomeDebouncePromise from 'awesome-debounce-promise'
import { HeadingXXLarge } from 'baseui/typography'
import { Head } from './Head'

const PAGE_INFO_FIELDS = gql`
  fragment PageInfoFields on PageInfo {
    hasNextPage
    endCursor
    pageCursor
  }
`

const GET_SEARCH_RESULTS = gql`
  ${ITEM_CARD_LIST_NODE_FIELDS}
  ${PAGE_INFO_FIELDS}
  query search($query: String!, $after: Cursor) {
    items: search(query: $query, after: $after, first: 10) {
      pageInfo {
        ...PageInfoFields
      }
      ...ItemCardListNodeFields
    }
  }
`

interface SearchResultsQueryData {
  items: ConnectionT<Story>
}

interface SearchResultsQueryVars {
  query: string
  after?: string
}

const SearchPage: React.FC = () => {
  const [searchParams, setSearchParams] = useSearchParams()
  const [queryTxt, setQueryTxt] = useState(searchParams.get('q') || '')
  const [isLoading, setIsLoading] = useState(false)
  const { loading, error, data, fetchMore, refetch } = useQuery<
    SearchResultsQueryData,
    SearchResultsQueryVars
  >(GET_SEARCH_RESULTS, {
    variables: {
      query: queryTxt,
    },
    notifyOnNetworkStatusChange: true,
    refetchWritePolicy: 'overwrite',
  })

  const [results, setResults] = useState<SearchResultsQueryData | undefined>(
    data
  )

  useEffect(() => {
    setResults(data)
  }, [data])

  useEffect(() => {
    setIsLoading(loading)
  }, [loading])

  const debounced = useCallback(
    AwesomeDebouncePromise((text: string) => {
      setResults(undefined)
      refetch({
        query: text,
      })
    }, 500),
    [queryTxt]
  )

  // input states
  const onQueryUpdate = useCallback(
    (val: string) => {
      setQueryTxt(val)
      debounced(val)
    },
    [setQueryTxt]
  )

  const onLoadMore = useCallback(() => {
    setIsLoading(true)
    fetchMore({
      query: GET_SEARCH_RESULTS,
      variables: {
        query: queryTxt,
        after: data?.items.pageInfo.pageCursor,
      },
    }).finally(() => {
      setIsLoading(false)
    })
  }, [queryTxt, data?.items.pageInfo.pageCursor, fetchMore])

  useEffect(() => {
    const params = new URLSearchParams()
    params.set('q', queryTxt)
    setSearchParams(params)
  }, [queryTxt])

  let children: React.ReactNode = <Fragment />
  if (!results && isLoading) {
    children = <LoadingScreen />
  } else if (!results && error) {
    children = <ErrorScreen error={error} />
  } else if (results) {
    children = (
      <Fragment>
        <Head />
        <SearchResults
          onLoadMore={onLoadMore}
          loading={isLoading}
          results={results.items}
        />
      </Fragment>
    )
  }
  return (
    <Container
      left={<Block />}
      center={
        <PaddedBlock>
          <SearchBar value={queryTxt} onChange={onQueryUpdate} />
          {children}
        </PaddedBlock>
      }
      right={<Block />}
    />
  )
}

interface SearchBarProps {
  value: string
  onChange: (val: string) => void
}

const SearchBar = ({ value, onChange }: SearchBarProps) => {
  const [css, theme] = useStyletron()
  const focusOut = useCallback((e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const el = document.getElementById('search-input')
    if (el) {
      el.blur()
    }
    return false
  }, [])
  return (
    <Block
      className={css({
        textAlign: 'center',
        paddingTop: theme.sizing.scale300,
      })}
    >
      <HeadingXXLarge
        className={css({
          textAlign: 'center',
          paddingBottom: theme.sizing.scale600,
        })}
      >
        Search HackerNews!
      </HeadingXXLarge>
      <form onSubmit={focusOut} onReset={focusOut}>
        <FormControl
          //label="Search HackerNews"
          caption="All search results are sorted by popularity."
        >
          <BaseInput
            id="search-input"
            name="q"
            placeholder="Type to search..."
            value={value}
            onChange={(e) => onChange(e.currentTarget.value)}
            autoFocus={true}
            overrides={{
              InputContainer: {
                style: ({ $theme }) => ({
                  border: `2px solid  ${$theme.colors.accent}`,
                }),
              },
            }}
          />
        </FormControl>
      </form>
    </Block>
  )
}

interface SearchResultsProps {
  onLoadMore: () => void
  loading: boolean
  results: ConnectionT<Story>
}

const SearchResults = ({
  onLoadMore,
  loading,
  results,
}: SearchResultsProps) => {
  return (
    <PaginatedItemCardList
      items={results}
      loading={loading}
      loadNext={onLoadMore}
      nextPageUrl={window.location.toString()}
    />
  )
}

export { SearchPage }
