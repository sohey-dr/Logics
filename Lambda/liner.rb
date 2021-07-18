require 'json'
require 'net/http'
require 'uri'

def lambda_handler(event:, context:)
    post_message(event["events"][0])

    { statusCode: 200, body: JSON.generate('sucess') }
end

def post_message(event)
    uri = URI.parse("https://api.line.me/v2/bot/message/reply")
    request = Net::HTTP::Post.new(uri)
    request.content_type = "application/json"
    request["Authorization"] = "Bearer #{ENV["CHANNEL_ACCESS_TOKEN"]}"
    request.body = json_dump(event)
    
    req_options = {
      use_ssl: uri.scheme == "https",
    }
    
    response = Net::HTTP.start(uri.hostname, uri.port, req_options) do |http|
      http.request(request)
    end
end

def json_dump(event)
    JSON.dump({
      "replyToken" => event["replyToken"],
      "messages" => [
        {
          "type" => "text",
          "text" => "#{find_hwid(event["beacon"]["hwid"])}\n
          ここのメッセージは自由に決めることができます！"
        }
      ]
    })
end

def find_hwid(hwid)
    case hwid
    when ENV["BEACON1"]
        "microbit1ですね！"
    when ENV["BEACON2"]
        "microbit2ですね！"
    end
end