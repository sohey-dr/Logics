(async() => {
  const Register = require("./libs/Register")
  const register = new Register()
  await register.run()
})();