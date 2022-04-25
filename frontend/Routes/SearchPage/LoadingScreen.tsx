import { Skeleton } from 'baseui/skeleton'
import { useStyletron } from 'baseui'
import { Fragment } from 'react'

const LoadingScreen: React.FC = () => {
  const [, theme] = useStyletron()
  const overrides = {
    Root: {
      style: {
        width: '100%',
        height: theme.sizing.scale4800,
        marginTop: theme.sizing.scale600,
        marginBottom: theme.sizing.scale900,
        borderRadius: theme.borders.radius300,
      },
    },
  }
  return (
    <Fragment>
      <Skeleton animation overrides={overrides} />
      <Skeleton animation overrides={overrides} />
    </Fragment>
  )
}

export { LoadingScreen }
