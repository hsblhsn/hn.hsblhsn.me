import { Block } from 'baseui/block'
import { gql, useQuery } from '@apollo/client'
import { Container } from '../../Components/Layout'
import { Item, ITEM_FIELDS } from '../../Components/Item'
import { useSearchParams } from 'react-router-dom'
import { PaddedBlock } from '../../Components/Layout'
import { LoadingScreen } from './LoadingScreen'
import { ErrorScreen } from './ErrorScreen'
import { Job, NodeT, Story } from '../../Types'
import { Fragment } from 'react'
import { Head } from './Head'

const GET_ITEM_QUERY = gql`
  ${ITEM_FIELDS}
  query GetItem($id: ID!) {
    item: node(id: $id) {
      ...ItemFields
      ...JobFields
    }
  }
`

interface GetItemQueryData {
  item: NodeT<Story | Job>
}

interface GetItemQueryVars {
  id: string
}

const ItemPage: React.FC = () => {
  const [searchParams] = useSearchParams()
  const { data, loading, error } = useQuery<GetItemQueryData, GetItemQueryVars>(
    GET_ITEM_QUERY,
    {
      variables: {
        id: searchParams.get('id') || '',
      },
      errorPolicy: 'all',
    }
  )
  let children: React.ReactNode = <Fragment />
  if (!data && loading) {
    children = <LoadingScreen />
  } else if (!data && error) {
    children = <ErrorScreen error={error} />
  } else if (data) {
    if (!data.item) {
      // there is an error or the item doesn't exist.
      children = (
        <Fragment>
          <Head item={data.item} /> <ErrorScreen error={error} />
        </Fragment>
      )
    } else {
      children = (
        <Fragment>
          <Head item={data.item} /> <Item item={data.item} />
        </Fragment>
      )
    }
  }
  return (
    <Container
      left={<Block />}
      center={<PaddedBlock>{children}</PaddedBlock>}
      right={<Block />}
    />
  )
}

export { ItemPage }
