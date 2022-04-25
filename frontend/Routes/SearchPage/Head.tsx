import config from '../../app.config'
import { getBestImage } from '../../Components/commonutils'
import { SEO } from '../../Components/SEO'
import { OpenGraph } from '../../Types'

interface HeadProps {
  item: {
    id: string
    title: string
    text?: string
    openGraph?: OpenGraph
  }
}

const Head: React.FC<HeadProps> = ({ item }: HeadProps) => {
  if (!item) {
    return (
      <SEO
        title={`404 Page Not Found | Hacker News`}
        description={'Page not found'}
        imageUrl={''}
        url={''}
      />
    )
  }
  const { id, title, text, openGraph } = item
  const description = openGraph?.description || text || ''
  return (
    <SEO
      title={`${openGraph?.title || title} | Hacker News`}
      description={description.replace(/<\/?[^>]+(>|$)/g, '')}
      imageUrl={
        getBestImage(openGraph?.image)?.url.replace(
          '&size=thumbnail',
          '&size=full'
        ) || ''
      }
      url={`${config.host}/item?id=${id}`}
    />
  )
}

export { Head }
