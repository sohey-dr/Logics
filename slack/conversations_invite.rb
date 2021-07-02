require 'net/http'
require 'uri'
require "json"
require 'pp'
require "csv"

=begin
CSVからプライベートチャンネルIDを取得していき、
Slack APIのconversations.invite(https://api.slack.com/methods/conversations.invite)を用いて
全てのチャンネルに招待するプログラム
=end

class ConversationsInvite
  attr_reader :token

  def initialize(token)
    TOKEN = token
  end

  def channel_id
    # これメモ
    @channel_id ||= CSV.read("slack/channels.csv").map{|row| row[1]}
  end

  def method_name
    
  end
end


ConversationsInvite.new().
