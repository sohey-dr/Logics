const fs = require("fs");
const csvSync = require("csv-parse/lib/sync");

module.exports = class CSVReader {
  constructor() {}

  run() {
    const users: Array<Array<string>> = this._readerCSV();
    return users;
  }

  _readerCSV() {
    const csv: Array<Array<string>> = fs.readFileSync("./perfect_user.csv");
    return csvSync(csv);
  }
};
