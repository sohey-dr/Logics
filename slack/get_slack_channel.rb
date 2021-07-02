require 'net/http'
require 'uri'
require "json"
require 'pp'
require "csv"

=begin
Slack APIのconversations.list(https://api.slack.com/methods/conversations.list)を用いてチャンネル一覧を取得するプログラム
メソッドを加えることで取得した値をCSVに書き出すようになっている
=end

class GetSlackChannel
  attr_reader :token

  def initialize(token)
    @token = token
  end

  def conversations_list
    http.use_ssl = uri.scheme === "https"
    write_csv
    puts "next cursor: #{next_cursor}" unless next_cursor.nil?
  end

  def write_csv
    channels.each do |c|
      channel_name = c["name"]
      channel_type = c["is_private"] ? find_private_channel_type(c) : "public"

      CSV.open('slack/channels.csv','a') do |csv|
        csv << [c["id"], channel_type, channel_name]
      end
    end
  end

  def find_private_channel_type(c)
    case 
    when /academy|career|boarding|transfer/ =~ c["name"]
      "user_private"
    when c["name"].include?("mentor-") || c["name"].include?("mentor_")
      "mentor_private"
    else
      "other_private"
    end
  end

  def channels
    response_hash["channels"]
  end

  def response_hash
    JSON.parse(response)
  end

  def response
    http.get(uri, headers).body
  end

  def req_url
    # cursorでページネーションを行う
    "https://slack.com/api/conversations.list?exclude_archived=true&limit=1000&types=public_channel,private_channel&pretty=1&cursor="
  end

  def uri
    @uri ||= URI.parse(req_url)
  end

  def http
    @http ||= Net::HTTP.new(uri.host, uri.port)
  end

  def headers
    { "Authorization" => "Bearer #{token}" }
  end

  def next_cursor
    # 1000件までしか取得できないため、次の情報を以下の値から取得できる
    response_hash["response_metadata"]["next_cursor"]
  end
end

GetSlackChannel.new("").conversations_list
