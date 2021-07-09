require "csv"

=begin
チャンネル一覧取得

=end
users.each do |user|
  user_name = user[5].gsub(" ", "")
end

class LookupPrivateChannel
  def 
    
  end

  def channels
    @channels ||= CSV.read("slack/tech_bowl/channels.csv")
  end

  def users
    @users ||= CSV.read("slack/tech_bowl/users(registered_slack_id).csv")
  end
end