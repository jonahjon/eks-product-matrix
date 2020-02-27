var fs = require('fs');
const tablemark = require('tablemark')

fs.readFile('./public/products.json', (err, data )=>{
    if (err) throw err;
    let products = JSON.parse(data);
    console.log(products.products);
    table = tablemark(products.products)
    fs.writeFile('TABLE.md', table, (err) => { 
        if (err) throw err; 
    }) 
})
  
  // | Name  | Age   | Is cool |
  // | ----- | ----- | ------- |
  // | Bob   | 21    | false   |
  // | Sarah | 22    | true    |
  // | Lee   | 23    | true    |
