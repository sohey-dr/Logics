require 'slack-ruby-client'
require "csv"

=begin
CSVからプライベートチャンネルIDを取得していき、
Slack APIのconversations.invite(https://api.slack.com/methods/conversations.invite)を用いて
全てのチャンネルに招待するプログラム
=end

Slack.configure do |conf|
  conf.token = "xoxp-324037009922-324037009986-2222883071830-9cb0c05866427228f477b5691c860aed"
end

client = Slack::Web::Client.new
CSV.foreach('slack/tech_bowl/channels.csv') do |row|
  if row[2] == "mentor_private"
    puts client.conversations_invite(channel: row[1], users: "UU6MYT400,U01GX8B06KH,U01D0EPHJ81,U020ZD483AR,U0204URQ6DR,U021MR7M5LP,U025JTP4Z4Z,U01TCCDBLTG,U01PNM7NA8J")
    sleep 1
  end
end
