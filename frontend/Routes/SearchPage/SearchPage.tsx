import { Block } from 'baseui/block'
import { Container, PaddedBlock } from '../../Components/Layout'
import { useSearchParams } from 'react-router-dom'
import { LoadingScreen } from './LoadingScreen'
import { ErrorScreen } from './ErrorScreen'
import { Fragment, useCallback, useEffect, useState } from 'react'
import { ITEM_CARD_LIST_NODE_FIELDS } from '../../Components/ItemCardList'
import { gql, useQuery } from '@apollo/client'
import { ConnectionT, Story } from '../../Types'
import AwesomeDebouncePromise from 'awesome-debounce-promise'
import { Head } from './Head'
import { SearchBar, SearchResults } from '../../Components/Search'

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
  const [query, setQuery] = useState(searchParams.get('q') || '')
  const [isLoading, setIsLoading] = useState(false)
  const { loading, error, data, fetchMore, refetch } = useQuery<
    SearchResultsQueryData,
    SearchResultsQueryVars
  >(GET_SEARCH_RESULTS, {
    variables: {
      query: query,
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

  const debouncedSearch = useCallback(
    AwesomeDebouncePromise((text: string) => {
      setResults(undefined)
      refetch({
        query: text,
      })
    }, 1000),
    [setResults, refetch]
  )

  // input states
  const onQueryUpdate = useCallback(
    (val: string) => {
      setQuery(val)
      debouncedSearch(val)
    },
    [setQuery, debouncedSearch]
  )

  const onLoadMore = useCallback(() => {
    setIsLoading(true)
    fetchMore({
      query: GET_SEARCH_RESULTS,
      variables: {
        query: query,
        after: data?.items.pageInfo.pageCursor,
      },
    }).finally(() => {
      setIsLoading(false)
    })
  }, [query, data?.items.pageInfo.pageCursor, fetchMore])

  useEffect(() => {
    const params = new URLSearchParams()
    params.set('q', query)
    setSearchParams(params)
  }, [query])

  let children: React.ReactNode = <Fragment />
  if (!results && isLoading) {
    children = <LoadingScreen />
  } else if (!results && error) {
    children = <ErrorScreen error={error} />
  } else if (results) {
    children = (
      <Fragment>
        <Head query={query} />
        <SearchResults
          query={query}
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
          <SearchBar value={query} onChange={onQueryUpdate} />
          {children}
        </PaddedBlock>
      }
      right={<Block />}
    />
  )
}

export { SearchPage }
