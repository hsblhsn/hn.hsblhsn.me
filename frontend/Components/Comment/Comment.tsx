import { Block } from 'baseui/block'
import { LabelSmall } from 'baseui/typography'
import { useStyletron } from 'baseui'
import { CommentThread } from '../CommentThread'
import { useState } from 'react'
import './Comment.css'
import { Comment as CommentT } from '../../Types'
import { StyleProps } from '../../Types'
import { fromNow } from '../commonutils'

interface CommentProps {
  comment: CommentT
}

const Comment: React.FC<CommentProps> = ({ comment }: CommentProps) => {
  const [css, theme] = useStyletron()
  const [isExpanded, setIsExpanded] = useState(true)
  const color = (c: string) =>
    css({
      color: c,
      cursor: 'pointer',
    })

  if (comment.dead || comment.deleted || comment.type !== 'comment') {
    // eslint-disable-next-line unicorn/no-null
    return null
  }

  return (
    <Block>
      <LabelSmall
        onClick={() => {
          setIsExpanded(!isExpanded)
        }}
        overrides={{
          Block: {
            style({ $theme }: StyleProps) {
              return {
                marginTop: $theme.sizing.scale300,
                marginBottom: $theme.sizing.scale300,
                paddingTop: $theme.sizing.scale300,
                paddingBottom: $theme.sizing.scale300,
                paddingLeft: $theme.sizing.scale300,
                borderRadius: $theme.borders.radius300,
                cursor: 'pointer',
                backgroundColor: isExpanded
                  ? $theme.colors.backgroundSecondary
                  : $theme.colors.backgroundTertiary,
                ':hover': {
                  opacity: '0.7',
                },
              }
            },
          },
        }}
      >
        <span className={color(theme.colors.accent)}>@{comment.by}</span>
        <span className={color(theme.colors.contentSecondary)}>
          &nbsp;commented&nbsp;
          {comment.time ? fromNow(comment.time * 1000) : 'unknown'}
        </span>
        <span className={color(theme.colors.contentSecondary)}>
          &nbsp;&nbsp;[{isExpanded ? 'close' : 'show'}]&nbsp;&nbsp;
        </span>
      </LabelSmall>
      <Block display={isExpanded ? 'block' : 'none'}>
        <div
          className="comment-reader-view-content"
          dangerouslySetInnerHTML={{ __html: comment.text || '' }}
        />
        <Block className="pl-3">
          <CommentThread
            parentId={comment.id}
            isChild={true}
            comments={comment.comments}
          />
        </Block>
      </Block>
    </Block>
  )
}

export { Comment }
export type { CommentProps }
