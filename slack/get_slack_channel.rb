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
    TOKEN = token
  end

  def channels
    response_hash["channels"]
  end

  def response_hash
    JSON.parse(response)
  end

  def response
    http.request(req)
  end

  def req_url
    "https://slack.com/api/conversations.list?types=public_channel,private_channel"
  end

  def uri
    @uri ||= URI.parse(req_url)
  end

  def http
    Net::HTTP.new(uri.host, uri.port)
  end

  def headers
    { "Authorization" => TOKEN }
  end

  def request
    Net::HTTP::Get.new(uri.path)
  end

  def set_header
    req.initialize_http_header(headers)
  end

  def next_cursor
    # 1000件までしか取得できないため、次の情報を以下の値から取得できる
    response_hash["response_metadata"]["next_cursor"]
  end
end

GetSlackChannel.new().