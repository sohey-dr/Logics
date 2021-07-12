require 'slack-ruby-client'
require "csv"

=begin
CSVからプライベートチャンネルIDを取得していき、
Slack APIのconversations.invite(https://api.slack.com/methods/conversations.invite)を用いて
全てのチャンネルに招待するプログラム
=end

Slack.configure do |conf|
  conf.token = ""
end

client = Slack::Web::Client.new
CSV.foreach('slack/not_have_channel.csv') do |row|
  
  client.conversations_create()
  # puts client.conversations_invite(channel: , users: "UU6MYT400,U01GX8B06KH,U01D0EPHJ81,U020ZD483AR,U0204URQ6DR,U021MR7M5LP,U025JTP4Z4Z,U01TCCDBLTG,U01PNM7NA8J,#{row[7]}")
end