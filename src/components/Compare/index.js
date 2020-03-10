import React, { Fragment } from 'react'
import 'font-awesome/css/font-awesome.min.css';
import './styles.css'
import '../../assets/styles.css'


String.prototype.Capitalize = function() {
  return this.charAt(0).toUpperCase() + this.slice(1);
}

function comparemaps(unique, arr1, arr2){
  var result1 = arr1.filter(function (o1) {
    return unique.some(function (o2) {
      if (o1.feature === o2.feature) {
        return (o1, o2);
      } else {
        return null;
      }
    });
  });

  var result2 = arr2.filter(function (o1) {
    return unique.some(function (o2) {
      if (o1.feature === o2.feature) {
        return (o1, o2);
      } else {
        return null;
      }
    });
  });
  var returnboth = 
  <Fragment>
    <td>
      <div class="status-badge">
        <div class="badge">
          <span class="icon">
            <i class="fa fa-check-circle"></i>
          </span>
        </div>
      </div>
    </td>
    <td>
      <div class="status-badge">
        <div class="badge">
          <span class="icon">
            <i class="fa fa-check-circle"></i>
          </span>
        </div>
      </div>
    </td>
  </Fragment>

  var returnthisfirst = 
    <Fragment>
    <td>
      <div class="status-badge">
        <div class="badge">
          <span class="icon">
            <i class="fa fa-check-circle"></i>
          </span>
        </div>
      </div>
    </td>
    <td>
    </td>
    </Fragment>

  var returnthissecond = 
    <Fragment>
      <td>
      </td>
      <td>
      <div class="status-badge">
        <div class="badge">
          <span class="icon">
            <i class="fa fa-check-circle"></i>
          </span>
        </div>
      </div>
    </td>
    </Fragment>

  if (result1.length === 0) {
    return returnthissecond
    } else if (result2.length === 0) {
      return returnthisfirst
    } else {
      return returnboth
    }
}

function createunique(d1) {
  var arr1 = d1[0].features
  var arr2 = d1[1].features
  arr1 = arr1.concat(arr2) // merge two arrays
  let foo = new Map();
  for(const tag of arr1) {
    foo.set(tag.feature, tag);
  }
  let uniquearray = [...foo.values()]
  return uniquearray
}


const Compare = ({ products }) =>
  <div className="row compare">
    <div className="col-12 mt-5 text-center">
      <table className="table">
        <thead className="thead-default">
          <tr>
            <th />
            {products.map(product =>
              <th key={product.id}>
                {product.name}
              </th>
            )}
          </tr>
        </thead>
        <tbody>
        <tr className="condition">
            <th scope="row">Type</th>
            {products.map(product =>
              <td key={product.id} className={product.saas_type === "Monitoring" ? "bg-blue" : "bg-green"}>
                {product.saas_type}
              </td>
            )}
          </tr>
          {createunique(products).map((features, index) => {
                return (
                  <Fragment>
                  <tr className="features">
                    <th scope="row">{features.feature}</th>
                      {comparemaps([features], products[0].features, products[1].features)}
                    </tr>
                  </Fragment>
                );
            })}
          <tr className="condition">
            <th scope="row">Price</th>
            {products.map(product =>
              <td key={product.id} className={"bg-red"}>
                {product.price}
              </td>
            )}
          </tr>
        </tbody>
      </table>
    </div>
  </div>;

export default Compare
