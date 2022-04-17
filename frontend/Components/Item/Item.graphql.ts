import { gql } from '@apollo/client'
import { COMMENT_THREAD_FIELDS } from '../CommentThread'

const ITEM_FIELDS = gql`
  ${COMMENT_THREAD_FIELDS}
  fragment ItemFields on Story {
    id
    by
    time
    url
    title
    text
    html
    type
    descendants
    comments(first: 5) {
      ...CommentThreadFields
    }
    openGraph {
      id
      title
      description
      url
      image {
        url
      }
    }
  }
  fragment JobFields on Job {
    id
    by
    time
    url
    title
    text
    html
    type
    openGraph {
      id
      title
      description
      url
      image {
        url
      }
    }
  }
`

export { ITEM_FIELDS }
