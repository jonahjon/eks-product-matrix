import React, { Fragment } from 'react'
import Badge from 'react-bootstrap/Badge'
import './styles.css'
import { compileFunction } from 'vm';

String.prototype.Capitalize = function() {
  return this.charAt(0).toUpperCase() + this.slice(1);
}

function whichlonger(d) {
  var longerarray = []
  var shorterarray = []
  if (d[0].features.length <= d[1].features.length) {
     longerarray = d[1].features;
     shorterarray = d[0].features;
  } else {
     longerarray = d[0].features;
     shorterarray = d[1].features;
  }
  return [longerarray, shorterarray]
}

function createunique(d) {
  var uniquearray = d[0].features.concat(d[1].features);
  var tempSet = new Set(uniquearray);
  uniquearray = Array.from(tempSet);
  console.log("uniquearray " + uniquearray)
  return uniquearray
}

function compare(arr1,arr2){
  const finalarray=[];
  arr1.forEach((e1)=>arr2.forEach((e2)=> {if(e1 === e2){
        finalarray.push(e1)
      }
    }
  ));
  return finalarray
}

var intersection = []

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
              <td key={product.id} className={product.type === "Monitoring" ? "bg-blue" : "bg-green"}>
                {product.type}
              </td>
            )}
          </tr>
          {console.log(compare(products[0].features, products[1].features))}
          {createunique(products).map((itemcount, index) => {
                return (
                  <Fragment>
                  <tr className="features">
                    <th scope="row">{itemcount.Capitalize()}</th>
                          {(index +1) <= (whichlonger(products))[1].length ?
                          <Fragment>
                            <td>
                              <Badge pill variant="success"></Badge>
                            </td>
                            <td>
                              <Badge pill variant="success"></Badge>
                            </td>
                          </Fragment>
                          :
                          <Fragment>
                            <td>
                              <Badge pill variant="success"></Badge>
                            </td>
                            <td>
                            </td>
                          </Fragment>
                      }
                    </tr>
                  </Fragment>
                );
            })}
          <tr className="condition">
            <th scope="row">Price</th>
            {products.map(product =>
              <td key={product.id} className={"bg-green"}>
                {product.price}
              </td>
            )}
          </tr>
        </tbody>
      </table>
    </div>
  </div>;

export default Compare
