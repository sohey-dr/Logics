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

users = ""
index = 0
CSV.foreach("hogehoge.csv") do |row|
  p row[0]
  begin
    puts client.conversations_invite(channel: "", users: row[0])
  # 既に招待されている場合はエラーが発生するので、無視する
  rescue Slack::Web::Api::Errors::AlreadyInChannel
    puts "次"
    next
  ensure
    sleep 1
  end
end
