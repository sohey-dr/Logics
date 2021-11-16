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
CSV.foreach('channels.csv') do |row|
  if row[2] == "mentor_private"
    begin
      puts client.conversations_invite(channel: row[1], users: "UU6MYT400")
    rescue Slack::Web::Api::Errors::AlreadyInChannel
      puts "次"
      next
    ensure
      sleep 1
    end
  end
end
