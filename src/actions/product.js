import * as types from '../constants/types'
import { getProductsQuery } from '../graphql/queries';
// import { graphql } from 'react-apollo';

// export default graphql(getProductsQuery)(Product);

const query = `
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

const url = "http://localhost:8080/graphql"

const opts = {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({ query })
};

export const getProductsgraph = () =>
  dispatch =>
    fetch(url, opts)
    .then(res => res.json())
    .then(res => {
      dispatch({
        type: types.FETCH_PRODUCTS,
        payload: res.data.getAllProducts
      })
    })

export const getProducts = () =>
  dispatch =>
    fetch(`products.json`)
      .then(response => response.json())
      .then(response => {
        dispatch({
          type: types.FETCH_PRODUCTS,
          payload: response.products
        })
      })


export const compare = product => ({
    type: types.COMPARE_PRODUCT,
    product
  })
