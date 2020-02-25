import { gql } from 'apollo-boost';


// example of a graphql mutation
const createProducts = gql`
  mutation createProducts($name: String!, $description: String) {
    createProducts(name: $name, description: $description) {
      name
      description
    }
  }
`;

export { createProducts };
