import { ErrorScreen } from '../../Components/ErrorScreen'
import { Container, PaddedBlock } from '../../Components/Layout'
import { Fragment } from 'react'

const NotFoundPage: React.FC = () => {
  return (
    <Container
      left={<Fragment />}
      center={
        <PaddedBlock>
          <ErrorScreen
            error={{
              title: '404',
              message: 'Page not found',
            }}
          />
        </PaddedBlock>
      }
      right={<Fragment />}
    />
  )
}

export { NotFoundPage }
