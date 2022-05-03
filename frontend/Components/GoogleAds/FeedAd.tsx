import React, { useEffect } from 'react'
import { Block } from 'baseui/block'
import { useStyletron } from 'baseui'

let googleAdTicker = 0

interface FeedAdProps {
  layoutKey?: string
  client: string
  slot: string
  testMode?: boolean
}

interface AdWindow extends Window {
  adsbygoogle?: Array<unknown>
}

const FeedAd: React.FC<FeedAdProps> = (props: FeedAdProps) => {
  const [css, theme] = useStyletron()
  useEffect(() => {
    const pushAd = () => {
      try {
        const adWindow: AdWindow = window
        const adsbygoogle = adWindow.adsbygoogle || []
        adsbygoogle.push({})
      } catch (e) {
        console.error(e)
      }
    }
    const interval = setInterval(() => {
      console.log('checking for GoogleAds')
      if (googleAdTicker > 15) {
        clearInterval(interval)
      }
      googleAdTicker++
      const adWindow: AdWindow = window
      // Check if Adsense script is loaded every 300ms
      if (adWindow.adsbygoogle) {
        pushAd()
        // clear the interval once the ad is pushed so that function isn't called indefinitely
        clearInterval(interval)
      }
    }, 300)

    return () => {
      clearInterval(interval)
    }
  }, [])
  return (
    <Block
      className={css({
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        border: `2px solid ${theme.colors.borderOpaque}`,
        marginBottom: theme.sizing.scale900,
        padding: theme.sizing.scale300,
        borderRadius: theme.borders.radius400,
        height: '200px',
        width: '100%',
        overflow: 'hidden',
        textDecoration: 'none',
      })}
    >
      <div
        className={css({
          position: 'absolute',
          zIndex: '1',
        })}
      >
        Imagine this is an ad!
      </div>
      <ins
        className="adsbygoogle"
        style={{
          display: 'block',
          zIndex: '5',
          height: '100px',
          width: '100%',
        }}
        data-ad-format="fluid"
        data-ad-client={props.client}
        data-ad-slot={props.slot}
        data-ad-layout-key={props.layoutKey || ''}
        data-adtest={props.testMode ? 'on' : 'off'}
      ></ins>
    </Block>
  )
}

export { FeedAd }
export type { FeedAdProps }
