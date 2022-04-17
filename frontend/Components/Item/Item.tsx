import { useStyletron } from 'baseui'
import { Block } from 'baseui/block'
import { Button } from 'baseui/button'
import { Cell, Grid } from 'baseui/layout-grid'
import { StyledLink } from 'baseui/link'
import {
  HeadingLarge,
  HeadingSmall,
  LabelXSmall,
  ParagraphXSmall,
} from 'baseui/typography'
import { useLocation } from 'react-router-dom'
import { Job, NodeT, Story } from '../../Types'
import { CommentThread } from '../CommentThread'
import { fromNow, getHost, getLink, getTitle } from '../commonutils'
import { ChevronLeft, ChevronRight } from 'baseui/icon'
import './Item.css'
import { useEffect } from 'react'

interface ItemProps {
  item: NodeT<Story | Job>
}

const Item: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()
  const location = useLocation()
  useEffect(() => {
    if (location.hash === '#comments') {
      document.getElementById('comments')?.scrollIntoView()
    }
  }, [location.key])

  if (!item || (item.type !== 'story' && item.type !== 'job')) {
    // eslint-disable-next-line unicorn/no-null
    return null
  }
  return (
    <Block paddingTop={theme.sizing.scale600}>
      <Header item={item} />
      <Content item={item} />
      <ActionButtons item={item} />
      <Comments item={item} />
      <ContentLinks item={item} />
    </Block>
  )
}

const Header: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [css, theme] = useStyletron()
  const color = (c: string) =>
    css({
      color: c,
      cursor: 'pointer',
    })
  return (
    <Block>
      <LabelXSmall>
        <span className={color(theme.colors.accent)}>@{item.by}</span>&nbsp;
        <span className={color(theme.colors.contentSecondary)}>
          {item.time ? fromNow(item.time * 1000) : 'unknown'}
        </span>
      </LabelXSmall>
      <HeadingLarge as="h1">
        {getTitle(item.title, item.openGraph?.title)}
      </HeadingLarge>
      <LabelXSmall
        color={theme.colors.contentTertiary}
        paddingBottom={theme.sizing.scale800}
      >
        Read on&nbsp;
        <StyledLink
          href={getLink(item.id, item.url)}
          target="_blank"
          rel="noreferrer"
        >
          {getHost(item.id, item.url)}
        </StyledLink>
      </LabelXSmall>
    </Block>
  )
}

const Content: React.FC<ItemProps> = ({ item }: ItemProps) => {
  let val = ''
  if ('text' in item && item.text !== '') {
    val = item.text
  } else if ('html' in item && item.html !== '') {
    val = item.html || ''
  } else if (item.openGraph?.description) {
    val = item.openGraph.description
  }
  return (
    <section
      id="reader-view-content"
      dangerouslySetInnerHTML={{ __html: val }}
    />
  )
}

const ActionButtons: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()
  const btnOverrides = {
    BaseButton: {
      style: {
        width: '100%',
      },
    },
  }
  const back = () => {
    window.history.back()
  }
  const open = () => {
    window.open(getLink(item.id, item.url), '_blank')
  }
  return (
    <Block
      paddingTop={theme.sizing.scale1200}
      paddingBottom={theme.sizing.scale600}
    >
      <Grid gridColumns={12} gridGaps={0} gridMargins={0}>
        <Cell span={6}>
          <Button
            onClick={back}
            kind="secondary"
            overrides={btnOverrides}
            startEnhancer={<ChevronLeft />}
          >
            Back
          </Button>
        </Cell>
        <Cell span={6}>
          <Button
            onClick={open}
            kind="secondary"
            overrides={btnOverrides}
            endEnhancer={<ChevronRight />}
          >
            Open
          </Button>
        </Cell>
      </Grid>
    </Block>
  )
}

const Comments: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()
  return (
    <Block id="comments">
      {'descendants' in item && (
        <HeadingSmall paddingTop={theme.sizing.scale600}>
          {item.descendants} comments
          <LabelXSmall color={theme.colors.contentTertiary}>
            Posted on&nbsp;
            <StyledLink
              href={getLink(item.id, undefined)}
              target="_blank"
              rel="noreferrer"
            >
              {getHost(item.id, undefined)}
            </StyledLink>
          </LabelXSmall>
        </HeadingSmall>
      )}
      {'comments' in item && (
        <Block>
          <CommentThread parentId={item.id} comments={item.comments} />
        </Block>
      )}
    </Block>
  )
}

const ContentLinks: React.FC<ItemProps> = ({ item }: ItemProps) => {
  const [, theme] = useStyletron()
  return (
    <Block>
      <ParagraphXSmall as="div" paddingTop={theme.sizing.scale2400}>
        <LabelXSmall>
          Contents:
          <br />
          <StyledLink href={getLink(item.id, item.url)}>
            {getLink(item.id, item.url)}
          </StyledLink>
        </LabelXSmall>
        <br />
        <LabelXSmall>
          Comments:
          <br />
          <StyledLink href={getLink(item.id, undefined)}>
            {getLink(item.id, undefined)}
          </StyledLink>
        </LabelXSmall>
      </ParagraphXSmall>
    </Block>
  )
}

export { Item }
export type { ItemProps }
