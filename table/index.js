var fs = require('fs');
const tablemark = require('tablemark')

fs.readFile('./public/new_products.json', (err, data )=>{
    if (err) throw err;
    let products = JSON.parse(data);
    console.log(products.security);
    table = tablemark(products.security)
    fs.writeFile('TABLE.md', table + "\r\n", (err) => { 
        if (err) throw err; 
    }) 
    console.log(products.monitoring);
    table = tablemark(products.monitoring)
    fs.appendFile('TABLE.md', table, (err) => { 
        if (err) throw err; 
    }) 
})