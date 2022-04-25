import { Grid, Cell, BEHAVIOR } from 'baseui/layout-grid'
import { Fragment } from 'react'

type ContainerProps = {
  left: JSX.Element
  center: JSX.Element
  right: JSX.Element
}

export function Container(props: ContainerProps) {
  return (
    <Fragment>
      <Grid
        behavior={BEHAVIOR.fixed}
        gridColumns={[4, 8, 12]}
        gridMargins={[0, 0, 0]}
        gridGaps={0}
        gridMaxWidth={1500}
      >
        <Cell span={[0, 1, 3]}>{props.left}</Cell>
        <Cell span={[4, 6, 6]}>{props.center}</Cell>
        <Cell span={[0, 1, 3]}>{props.right}</Cell>
      </Grid>
    </Fragment>
  )
}
