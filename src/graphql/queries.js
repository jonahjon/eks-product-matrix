import { gql } from 'apollo-boost';

//graphql query
const getProductsQuery = gql`
    query {
        getAllProducts {
            name
            image
            saas_type
            price
            description
        }
    }
`;

export { getProductsQuery };
