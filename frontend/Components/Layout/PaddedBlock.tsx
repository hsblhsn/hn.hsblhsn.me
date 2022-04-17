import { styled } from 'baseui'
import { Block } from 'baseui/block'

const PaddedBlock = styled(Block, ({ $theme }) => ({
  paddingLeft: $theme.sizing.scale600,
  paddingRight: $theme.sizing.scale600,
  paddingTop: $theme.sizing.scale600,
  paddingBottom: $theme.sizing.scale300,
}))

export { PaddedBlock }
