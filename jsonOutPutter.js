const express = require('express')
const app = express()
app.use(express.json())
app.use(express.urlencoded({ extended: true }));

app.post('/', function (req, res) {
  console.log(req.body);
  res.send('Got a POST request')
})

app.listen(3000, () => console.log('Example app listening on port 3000!'))