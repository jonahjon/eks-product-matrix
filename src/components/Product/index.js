import React from 'react'
import './styles.css'
import Badge from 'react-bootstrap/Badge'

String.prototype.Capitalize = function() {
    return this.charAt(0).toUpperCase() + this.slice(1);
}

const Product = ({product, compare}) =>
    <div key={product.id} className="col-sm-6 col-md-3">
        <div className={"product " + (product.compare ? "compare" : "")} >
            <img src={product.image} alt={product.name} />
            <div className={product.saas_type === "Monitoring" ? "image_overlay_monitoring" : "image_overlay_security"}/>
            <div className="view_details" onClick={() => compare(product)}>
              {product.compare ? "Remove" : "Compare"}
            </div>
            <div className="stats">
                <div className="stats-container">
                    <span className="align-items-center">{product.name}</span>
                    <span className="product_badge">
                        <Badge pill variant={product.saas_type === "Monitoring" ? "primary" : "success"}>
                            {product.saas_type}
                        </Badge>
                    </span>
                    <span className="align-items-center"></span>
                    <p></p>
                    {product.features.map(function(feature, index) {
                            return(
                                <div>
                                    <span className="align-items-center">{feature.feature.Capitalize()}</span>
                                    <span className="product_badge"><i class="fa fa-check-circle"></i></span>
                                </div>
                            )
                        }
                        )}
                </div>
            </div>
        </div>
    </div>;

export default (Product);
