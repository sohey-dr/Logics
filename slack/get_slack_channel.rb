require 'net/http'
require 'uri'
require "json"
require 'pp'

TOKEN = ""
SLACK_API_BASE = 'https://slack.com/api'

fetch_all_channels_url = "https://slack.com/api/conversations.list?token=#{TOKEN}&types=public_channel,private_channel&limit=1000"
res = Net::HTTP.get(URI.parse(fetch_all_channels_url))
channel_hash = JSON.parse(res)
channels = channel_hash["channels"]
channels = channels.each { |channel| puts channel["name"] }

# 1000件までしか取得できないため、次の情報を以下の値から取得できる
puts channel_hash["response_metadata"]["next_cursor"]