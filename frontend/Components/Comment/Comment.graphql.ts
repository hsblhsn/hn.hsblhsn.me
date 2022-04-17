import { gql } from '@apollo/client'

const COMMENT_FIELDS = gql`
  fragment CommentContentFields on Comment {
    id
    type
    by
    time
    text
    dead
    deleted
  }
  fragment CommentFields on Comment {
    ...CommentContentFields
    comments(first: 5) {
      pageInfo {
        hasNextPage
        endCursor
      }
      edges {
        node {
          ...CommentContentFields
          comments(first: 5) {
            pageInfo {
              hasNextPage
              endCursor
            }
            edges {
              node {
                ...CommentContentFields
              }
            }
          }
        }
      }
    }
  }
`

export { COMMENT_FIELDS }
