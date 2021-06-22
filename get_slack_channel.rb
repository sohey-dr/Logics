require 'net/http'
require 'uri'
require "json"
require 'pp'

TOKEN = ""
SLACK_API_BASE = 'https://slack.com/api'


fetch_all_channels_url = "https://slack.com/api/conversations.list?token=#{TOKEN}"
res = Net::HTTP.get(URI.parse(fetch_all_channels_url))
channel_hash = JSON.parse(res)
channels = channel_hash["channels"]
channels = channels.each { |channel| p channel["name"] }