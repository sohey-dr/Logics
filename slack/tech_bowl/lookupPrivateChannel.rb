require "csv"

=begin
チャンネル一覧取得して
csvから登録されていないユーザーを一気に探す処理
既に登録してる
=end

class LookupPrivateChannel
  def write_csv
    CSV.open("slack/tech_bowl/perfect_user.csv", "w") do |csv|
      set_users.each do |user|
        csv << user
      end
    end
  end

  def set_users
    users.each do |user|
      user[6] = lookup_channel_by_user_name(user[5].gsub(" ", ""))
    end
  end

  def lookup_channel_by_user_name(user_name)
    channel_name = ""
    channels.each do |channel|
      if channel[3].include?(user_name)
        channel_name = channel[1]
        break
      end
    end
    channel_name == '' ? nil : channel_name
  end

  def channels
    @channels ||= CSV.read("slack/tech_bowl/channels.csv")
  end

  def users
    @users ||= CSV.read("slack/tech_bowl/not_have_channel.csv")
  end
end

LookupPrivateChannel.new.write_csv