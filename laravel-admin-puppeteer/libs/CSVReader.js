const fs = require('fs');
const csvSync = require('csv-parse/lib/sync');

module.exports = class CSVReader {
  constructor() {}

  run() {
    const users = this._readerCSV();
    return users;
  }

  _readerCSV() {
    const csv = fs.readFileSync('laravel-admin-puppeteer/perfect_user.csv');
    return csvSync(csv);
  }
}
