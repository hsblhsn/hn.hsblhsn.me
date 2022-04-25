import { Route, Routes as ReactRouterRoutes } from 'react-router-dom'
import { HomePage } from './HomePage'
import { ItemPage } from './ItemPage'
import { NotFoundPage } from './NotFoundPage'
import { SearchPage } from './SearchPage'

const Routes = () => (
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
)

export { Routes }
