require "sinatra"
require "json"

post "/", provides: :json do
  params = JSON.parse request.body.read
  p params
  return params["challenge"]
end
