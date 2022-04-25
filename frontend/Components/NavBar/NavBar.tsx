import { Block } from 'baseui/block'
import {
  StyledNavigationItem,
  StyledNavigationList,
  ALIGN,
} from 'baseui/header-navigation'
import { Link } from 'react-router-dom'
import { useStyletron } from 'baseui'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface NavBarProps {}

const NavBar: React.FC<NavBarProps> = () => {
  const [, theme] = useStyletron()
  return (
    <Block>
      <Block
        $style={{
          backgroundColor: '#ff6600',
          display: 'flex',
          paddingBottom: theme.sizing.scale500,
          paddingTop: theme.sizing.scale500,
          borderBottomWidth: '1px',
          borderBottomStyle: 'solid',
          borderBottomColor: `${theme.colors.borderOpaque}`,
        }}
      >
        <StyledNavigationList $align={ALIGN.left}>
          <StyledNavigationItem className="hidden md:block">
            <Link to="/"> HackerNews</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/">Top</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/newest">New</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/ask">Ask</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/show">Show</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/jobs">Jobs</Link>
          </StyledNavigationItem>
          <StyledNavigationItem>
            <Link to="/search">Search</Link>
          </StyledNavigationItem>
        </StyledNavigationList>
      </Block>
    </Block>
  )
}

export { NavBar }
export type { NavBarProps }
