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
CSV.foreach('slack/tech_bowl/channels.csv') do |row|
  if row[2] == "mentor_private"
    puts client.conversations_invite(channel: row[1], users: "UU6MYT400")
  end
  # sleep 1
end
