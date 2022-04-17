import { Button, KIND } from 'baseui/button'
import { ConnectionT, Job, Story, StyleProps } from '../../Types'
import { ItemCardList } from './ItemCardList'

interface PaginatedItemCardListProps {
  items: ConnectionT<Story | Job>
  loading: boolean
  loadNext: () => void
  nextPageUrl: string
}

const PaginatedItemCardList: React.FC<PaginatedItemCardListProps> = ({
  items,
  loading,
  loadNext,
  nextPageUrl,
}: PaginatedItemCardListProps) => {
  const linkToNextPage = (
    <a
      href={nextPageUrl}
      onClick={(e) => {
        e.preventDefault()
        return false
      }}
    >
      Load more
    </a>
  )
  return (
    <div>
      <ItemCardList items={items} />
      <Button
        kind={KIND.secondary}
        onClick={loadNext}
        isLoading={loading}
        disabled={loading || items.pageInfo?.hasNextPage === false}
        overrides={{
          Root: {
            style: ({ $theme }: StyleProps) => ({
              width: '100%',
              marginTop: $theme.sizing.scale900,
            }),
          },
        }}
      >
        {items.pageInfo?.hasNextPage ? linkToNextPage : 'No more items'}
      </Button>
    </div>
  )
}

export { PaginatedItemCardList }
export type { PaginatedItemCardListProps }
