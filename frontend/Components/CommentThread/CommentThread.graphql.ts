import { gql } from '@apollo/client'
import { COMMENT_FIELDS } from '../Comment'

const COMMENT_THREAD_FIELDS = gql`
  ${COMMENT_FIELDS}
  fragment CommentThreadFields on CommentConnection {
    pageInfo {
      hasNextPage
      endCursor
    }
    edges {
      node {
        ...CommentFields
      }
    }
  }
`

const LOAD_MORE_COMMENTS_QUERY = gql`
  ${COMMENT_THREAD_FIELDS}
  query LoadMoreComments($parentId: ID!, $after: Cursor) {
    item: node(id: $parentId) {
      ... on Comment {
        id
        comments(after: $after, first: 5) {
          ...CommentThreadFields
        }
      }
      ... on Story {
        id
        comments(after: $after, first: 5) {
          ...CommentThreadFields
        }
      }
    }
  }
`

export { COMMENT_THREAD_FIELDS, LOAD_MORE_COMMENTS_QUERY }
