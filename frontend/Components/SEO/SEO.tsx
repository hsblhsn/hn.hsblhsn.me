import { Helmet } from 'react-helmet-async'
import { useContext } from 'react'
import { ConfigContext } from '../Config'

interface SEOProps {
  title: string
  description: string
  url: string
  imageUrl?: string
}

const SEO: React.FC<SEOProps> = (props: SEOProps) => {
  const config = useContext(ConfigContext)
  let imageUrl = props.imageUrl
  if (props.title !== '') {
    let title = props.title
    const trimSuffix = '| Hacker News'
    if (props.title.endsWith(trimSuffix)) {
      title = props.title.substring(0, props.title.length - trimSuffix.length)
    }
    imageUrl = `${config.host}/images/social_preview.jpeg?title=${title}`
  }
  if (!imageUrl || imageUrl.length === 0) {
    imageUrl = `${config.host}/og-banner.jpg`
  }
  return (
    <Helmet>
      <title>{props.title}</title>
      <meta property="og:title" content={props.title} />
      <meta name="description" content={props.description} />
      <meta property="og:description" content={props.description} />
      <meta property="og:image" content={imageUrl} />
      <meta property="og:url" content={props.url} />
    </Helmet>
  )
}

export { SEO }
