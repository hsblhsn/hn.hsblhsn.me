import { useStyletron } from 'baseui'
import { Block } from 'baseui/block'
import { CommentConnection } from '../../Types'
import { Comment } from '../Comment'
import { ApolloError, useApolloClient } from '@apollo/client'
import { LOAD_MORE_COMMENTS_QUERY } from './CommentThread.graphql'
import React, { useState } from 'react'
import { Button, KIND, SIZE } from 'baseui/button'
import { ChevronDown } from 'baseui/icon'

interface CommentThreadProps {
  parentId: string
  comments: CommentConnection | undefined
  isChild?: boolean
}

const CommentThread: React.FC<CommentThreadProps> = ({
  parentId,
  comments,
  isChild = false,
}: CommentThreadProps) => {
  const [css, theme] = useStyletron()
  if (!comments) {
    // eslint-disable-next-line unicorn/no-null
    return null
  }

  const threadStyle = (child: boolean) => {
    return css({
      paddingLeft: child ? theme.sizing.scale300 : theme.sizing.scale0,
      borderLeft: child ? `3px solid ${theme.colors.accent}` : 'none',
    })
  }
  return (
    <Block className={threadStyle(isChild)}>
      {comments.edges.map(({ node }, index) => {
        if (!node) {
          // eslint-disable-next-line unicorn/no-null
          return null
        }
        return <Comment comment={node} key={index} />
      })}
      {comments.pageInfo.hasNextPage && (
        <LoadMoreComments
          isChild={isChild}
          parentId={parentId}
          after={comments.pageInfo.endCursor}
        />
      )}
    </Block>
  )
}

interface LoadMoreCommentsProps {
  parentId: string
  after: string
  isChild: boolean
}

const LoadMoreComments: React.FC<LoadMoreCommentsProps> = ({
  isChild,
  parentId,
  after,
}: LoadMoreCommentsProps) => {
  const client = useApolloClient()
  const [loading, setLoading] = useState(false)
  const [comments, setComments] = useState<CommentConnection | undefined>()
  const [error, setError] = useState<ApolloError | undefined>()

  const onLoadAfter = (parentId: string, after: string) => {
    setLoading(true)
    client
      .query({
        query: LOAD_MORE_COMMENTS_QUERY,
        variables: {
          parentId: parentId,
          after: after,
        },
        fetchPolicy: 'cache-first',
        errorPolicy: 'ignore',
      })
      .then(({ data, loading, error }) => {
        setLoading(loading)
        setComments(data.item.comments)
        setError(error)
      })
      .catch((error) => {
        setError(error)
      })
      .finally(() => {
        setLoading(false)
      })
  }
  if (comments) {
    return (
      <CommentThread
        parentId={parentId}
        comments={comments}
        isChild={isChild}
      />
    )
  }
  return (
    <Button
      isLoading={loading}
      disabled={loading || error !== undefined}
      kind={KIND.tertiary}
      size={SIZE.compact}
      startEnhancer={<ChevronDown size={24} />}
      overrides={{
        BaseButton: {
          style: ({ $theme }) => {
            return {
              marginTop: $theme.sizing.scale600,
              color: $theme.colors.accent,
            }
          },
        },
      }}
      onClick={() => {
        onLoadAfter(parentId, after)
      }}
    >
      {error ? 'Could not load' : 'Load more comments'}
    </Button>
  )
}

export { CommentThread }
export type { CommentThreadProps }
