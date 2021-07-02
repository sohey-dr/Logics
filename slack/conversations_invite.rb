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
CSV.foreach('slack/channels.csv') do |row|
  if row[1] == "user_private"
    puts client.conversations_invite(channel: row[0], users: "UU6MYT400,U01GX8B06KH,U01D0EPHJ81,U020ZD483AR,U0204URQ6DR,U021MR7M5LP,U025JTP4Z4Z,U01TCCDBLTG,U01PNM7NA8J")
  end
  # sleep 1
end
