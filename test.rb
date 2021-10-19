require "net/http"

uri = URI.parse("https://www.google.com/search?q=%E5%90%9B%E3%81%8C%E5%A5%BD%E3%81%8D+%E6%AD%8C%E8%A9%9E")
response = Net::HTTP.get_response(uri)

response.code # status code
body = response.body

File.open("result.html", mode = "w"){|f|
  f.write(body)
}