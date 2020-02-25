import { InMemoryCache } from 'apollo-cache-inmemory';
import ApolloClient from 'apollo-client/ApolloClient';
import { ApolloLink } from 'apollo-link';
import { setContext } from 'apollo-link-context';
import { onError } from 'apollo-link-error';
import { createHttpLink } from 'apollo-link-http';
import { getGraphQLEndpoint } from '../utils/config.js';

//const _log = getLogger('apollo');

const requestLink = createHttpLink({
    uri: getGraphQLEndpoint()
});

const httpLink = setContext((_, {headers}) => ({
    headers: {
        ...headers,
    }
}));

const errorLink = onError(({ graphQLErrors, networkError }) => {
  if (graphQLErrors)
    graphQLErrors.map(({ message, locations, path }) =>
      console.log(
        `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`,
      ),
    );
  if (networkError) console.log(`[Network error]: ${networkError}`);
});

// assemble and export the client
export const apolloClient = new ApolloClient({
    cache: new InMemoryCache(),
    link: ApolloLink.from([
        errorLink,
        httpLink,
        requestLink
    ]),
    defaultOptions: {}
});