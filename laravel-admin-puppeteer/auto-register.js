(async() => {
  const Register = require("./libs/Register")
  const CSVReader = require("./libs/CSVReader")
  for await (const user of users) {
    const register = new Register(user);
    await register.run();
  };
})();