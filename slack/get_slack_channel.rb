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

  def get_slack_channels
    http.use_ssl = uri.scheme === "https"
    write_csv
  end

  def write_csv
    channels.each do |c|
      channel_name = c["name"]
      channel_type = nil

      if /academy|career|boarding|transfer/ =~ channel_name
        channel_type = 0
      else
        channel_type = 1
      end

      puts "#{channel_type}, #{channel_name}"
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
    "https://slack.com/api/conversations.list?exclude_archived=true&limit=1000&types=private_channel&pretty=1"
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

GetSlackChannel.new("").get_slack_channels