import React, { Component } from 'react'
import {Product} from '../'
// import { getProductsQuery } from '../../graphql/queries';
// import { graphql } from 'react-apollo';
// export default graphql(getProductsQuery)(Product);

const ProductList = ({products, compare}) =>
  <div className="row mt-3">
      {products.map(product =>
        <Product key={product.id} product={product} compare={compare} />
      )}
  </div>;

export default ProductList
