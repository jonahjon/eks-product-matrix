import { gql } from 'apollo-boost';

//graphql query
const getProducts = gql`
   {
        getProducts {
            name
            description
        }
    }
`;

export { getProducts };
