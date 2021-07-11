const fs = require("fs");
const csvSync = require("csv-parse/lib/sync");

module.exports = class CSVReader {
  constructor() {}

  run() {
    const users = this._readerCSV();
    return users;
  }

  _readerCSV() {
    // csvのカラムごとに分けて多次元配列として返す
    const csv = fs.readFileSync("./perfect_user.csv");
    return csvSync(csv);
  }
};
