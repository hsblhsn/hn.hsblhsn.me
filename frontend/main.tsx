import React from 'react'
import ReactDOM from 'react-dom/client'
import { Client as Styletron } from 'styletron-engine-atomic'
import { Provider as StyletronProvider } from 'styletron-react'
import './index.css'
import { App } from './Components/App'
import { BrowserRouter } from 'react-router-dom'
import { HelmetProvider } from 'react-helmet-async'
import {
  ApolloProvider,
  ApolloClient,
  InMemoryCache,
  HttpLink,
} from '@apollo/client'
import { relayStylePagination } from '@apollo/client/utilities'
import { Routes } from './Routes'
import config from './app.config'

const engine = new Styletron()

const client = new ApolloClient({
  cache: new InMemoryCache({
    typePolicies: {
      Query: {
        fields: {
          newStories: relayStylePagination(),
          showStories: relayStylePagination(),
          topStories: relayStylePagination(),
          jobStories: relayStylePagination(),
          askStories: relayStylePagination(),
        },
      },
    },
  }),
  link: new HttpLink({
    uri: config.graphqlEndpoint,
    useGETForQueries: true,
  }),
})

const container = document.getElementById('root') as HTMLElement
const root = ReactDOM.createRoot(container)

root.render(
  <React.StrictMode>
    <StyletronProvider value={engine}>
      <ApolloProvider client={client}>
        <BrowserRouter>
          <HelmetProvider>
            <App>
              <Routes />
            </App>
          </HelmetProvider>
        </BrowserRouter>
      </ApolloProvider>
    </StyletronProvider>
  </React.StrictMode>
)
