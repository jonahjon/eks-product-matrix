import React from 'react'
import './styles.css'
import Badge from 'react-bootstrap/Badge'
import { getProductsQuery } from '../../graphql/queries';
import { graphql } from 'react-apollo';

String.prototype.Capitalize = function() {
    return this.charAt(0).toUpperCase() + this.slice(1);
}

const Product = ({product, compare}) =>
    <div key={product.id} className="col-sm-6 col-md-3">
        <div className={"product " + (product.compare ? "compare" : "")} >
            <img src={product.image} alt={product.name} />
            <div className="image_overlay"/>
            <div className="view_details" onClick={() => compare(product)}>
              {product.compare ? "Remove" : "Compare"}
            </div>
            <div className="stats">
                <div className="stats-container">
                    <span className="align-items-center">{product.name}</span>
                    <span className="product_badge">
                        <Badge pill variant="primary">
                            {product.saas_type}
                        </Badge>
                    </span>
                    <span className="align-items-center"></span>
                    <p></p>
                    {product.features.map(function(item, index) {
                            return(
                                <div>
                                    <span className="align-items-center">{item.Capitalize()}</span>
                                    <span className="product_badge"><Badge pill variant="success">Yes</Badge></span>
                                </div>
                            )
                        }
                        )}
                    <p>{product.description}</p>
                </div>
            </div>
        </div>
    </div>;

export default graphql(getProductsQuery)(Product);
