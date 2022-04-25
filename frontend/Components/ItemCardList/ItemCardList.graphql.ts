import { gql } from '@apollo/client'
import { ITEM_CARD_STORY_FIELDS, ITEM_CARD_JOB_FIELDS } from '../ItemCard'

const ITEM_CARD_LIST_NODE_FIELDS = gql`
  ${ITEM_CARD_STORY_FIELDS}
  fragment ItemCardListNodeFields on NodeConnection {
    edges {
      node {
        ...StoryCardFields
      }
    }
  }
`

const ITEM_CARD_LIST_STORY_FIELDS = gql`
  ${ITEM_CARD_STORY_FIELDS}
  fragment ItemCardListStoryFields on StoryConnection {
    edges {
      node {
        ...StoryCardFields
      }
    }
  }
`

const ITEM_CARD_LIST_JOB_FIELDS = gql`
  ${ITEM_CARD_JOB_FIELDS}
  fragment ItemCardListJobFields on JobConnection {
    edges {
      node {
        ...JobCardFields
      }
    }
  }
`

export {
  ITEM_CARD_LIST_NODE_FIELDS,
  ITEM_CARD_LIST_STORY_FIELDS,
  ITEM_CARD_LIST_JOB_FIELDS,
}
