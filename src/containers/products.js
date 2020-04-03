import productlist from './products.json'


export const productJSON = productlist.products.map(
  ({ name, saas_type }) => {
    return{
      value: saas_type,
      label: name
    }
  }
);