#### addProduct
```
mutation ($name: String!, $image: String!, $saas_type: String!, $price: String!, $description: String!) {
  addProduct(name: $name, image: $image, saas_type: $saas_type, price: $price, description: $description) {
    ...MinimalRecipe
  }
}

fragment MinimalRecipe on Product {
  name
  image
  saas_type
  price
  description
}

{
  "name": "DynaTrace",
  "image": "dynatrace.png",
  "saas_type": "infra",
  "price": "$5 / day",
  "description": "used to monitor stuff"
}

```

#### getProduct
```
query ($name: String!) {
  getProduct(name: $name) {
    ...SampleProduct
  }
}

fragment SampleProduct on Product {
  _id
  name
  image
  saas_type
  price
  description
  features{feature}
}

{
  "name": "Data Dog",
}
```


#### getAllProducts
```
{
  getAllProducts {
    ...AllProduct
  }
}



fragment AllProduct on Product {
  _id
  name
  image
  saas_type
  price
	description
  features {feature}
}

```




#### updateProduct
```
mutation ($name: String!, $image: String!, $saas_type: String!, $price: String!, $description: String!) {
  updateProduct(name: $name, image: $image, saas_type: $saas_type, price: $price, description: $description) {
    ...SampleProduct
  }
}

fragment SampleProduct on Product {
  _id
  name
  image
  saas_type
  price
  description
}

{
  "description": "used to monitor stuff",
  "image": "dynatrace.png",
  "name": "DynaTrace",
  "price": "$6 / day",
  "saas_type": "infra"
}
```

#### deleteProduct
```
mutation ($name: String!, $image: String!, $saas_type: String!, $price: String!, $description: String!) {
  updateProduct(name: $name, image: $image, saas_type: $saas_type, price: $price, description: $description) {
    ...SampleProduct
  }
}

fragment SampleProduct on Product {
  _id
  name
  image
  saas_type
  price
  description
}

{
  "description": "used to monitor stuff",
  "image": "dynatrace.png",
  "name": "DynaTrace",
  "price": "$6 / day",
  "saas_type": "infra"
}
```