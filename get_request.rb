require "net/http"

uri = URI.parse("https://k7qropickw.summer2021.dena.jp/channels/6834/messages")
response = Net::HTTP.get_response(uri)

response.code # status code
body = response.body
i = 0

body.each do |b|
  i += 1
  b
end
