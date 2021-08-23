require "sinatra"
require "json"

post "/", provides: :json do
  params = JSON.parse request.body.read
  p params
  # これはslackのevent api
  return params["challenge"]
end
