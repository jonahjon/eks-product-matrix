import registerServiceWorker from './registerServiceWorker'
import React from 'react'
import ReactDOM from 'react-dom'
import {BrowserRouter} from 'react-router-dom'
import {Provider} from 'react-redux'
import App from './containers/App'
import { reduxStore } from './redux/root-reducer';
import { apolloClient } from './graphql/apollo.client';
import { ApolloProvider } from 'react-apollo';
import 'bootstrap/dist/css/bootstrap.css'
import './styles.css'

ReactDOM.render(
  <Provider store={reduxStore}>
    <ApolloProvider client={apolloClient}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </ApolloProvider>
  </Provider>,
  document.getElementById('root')
);

registerServiceWorker();
