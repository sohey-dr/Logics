const fs = require('fs');
const csv = require('csv');

fs.createReadStream('laravel-admin-puppeteer/perfect_user.csv')
  .pipe(csv.parse(function(err, data) {
 
        console.log(data);
 
  }));
