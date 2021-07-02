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
user_list.members.each do |user|
  client.conversations_invite(channel: "G01E5PDMSNT", users: "U01PNM7NA8J")
end