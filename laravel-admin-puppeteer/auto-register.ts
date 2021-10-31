(async() => {
  const CSVReader = require("./libs/CSVReader")
  const csvReader = new CSVReader
  const users: Array<Array<string>> = csvReader.run();
  const Register = require("./libs/Register");
  for await (const user of users) {
    const register = new Register(user);
    await register.run();
  };
})();