(async() => {
  const Register = require("./libs/Register")
  const CSVReader = require("./libs/CSVReader")
  const users = CSVReader.run();
  users.forEach(user => {
    const register = new Register(user);
    await register.run();
  });
})();