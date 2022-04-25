import { Route, Routes as ReactRouterRoutes } from 'react-router-dom'
import HomePage from './HomePage'
import { lazy, Suspense, useEffect, useRef } from 'react'
import { useStyletron } from 'baseui'
import LoadingBar from 'react-top-loading-bar'

const ItemPage = lazy(() => import('./ItemPage'))
const SearchPage = lazy(() => import('./SearchPage'))
const NotFoundPage = lazy(() => import('./NotFoundPage'))

const Routes = () => (
  <Suspense fallback={<Fallback />}>
    <ReactRouterRoutes>
      <Route path="/" element={<HomePage />} />
      <Route path="/newest" element={<HomePage />} />
      <Route path="/ask" element={<HomePage />} />
      <Route path="/show" element={<HomePage />} />
      <Route path="/jobs" element={<HomePage />} />
      <Route path="/item" element={<ItemPage />} />
      <Route path="/search" element={<SearchPage />} />
      <Route path="*" element={<NotFoundPage />} />
    </ReactRouterRoutes>
  </Suspense>
)

interface LoadingBarRefI {
  staticStart(): void
  continuousStart(): void
  complete(): void
}

const Fallback = () => {
  const ref = useRef<LoadingBarRefI>(null)
  const [, theme] = useStyletron()
  useEffect(() => {
    ref.current?.continuousStart()
  }, [])
  return (
    <LoadingBar
      color={theme.colors.backgroundAlwaysDark}
      height={3}
      ref={ref}
      shadow={true}
    />
  )
}

export { Routes }
