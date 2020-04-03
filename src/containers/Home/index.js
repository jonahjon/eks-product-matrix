import React from 'react'
import {bindActionCreators} from 'redux'
import {Compare, ProductList} from '../../components'
import { Search } from '../../components/Search/search.jsx' 
import * as productActions from '../../actions/product'
import Select from 'react-select';
import { productJSON } from '../products.js'
import './index.css';

import {connect} from 'react-redux'

class Home extends React.Component {
  constructor() {
    super();
    this.state = {
      // productsJSON: fetch(`products.json`),
      searchfield: ''
    }
  }

  handleChange = (e) => {
    this.setState({searchfield: e.target.value});
  };

  componentWillMount() {
    this.props.actions.getProducts()
  }
  
  render() {
    const { searchfield } = this.state;
    const {products, actions} = this.props;
    const filteredData = products.filter(item => {
      return Object.keys(item).some(key =>
        typeof item[key] === "string" && item[key].toLowerCase().includes(searchfield.toLowerCase())
      );
    });
    const compareProducts = products.filter(product => product.compare);
    return (
      <div className="home mt-5">
        <h1>EKS Partner Matrix UX Demo</h1>
        <Search
            type='search'
            placeholder='Security'
            handleChange={this.handleChange}
         />
        <ProductList products={filteredData} compare={actions.compare}/>
        {compareProducts.length >= 2 &&
          <Compare products={compareProducts}/>
        }
      </div>
    )
  }
}

export default connect(
  state => ({
    products: state.product.products
  }),
  dispatch => ({
    actions: bindActionCreators(productActions, dispatch)
  })
)(Home);
