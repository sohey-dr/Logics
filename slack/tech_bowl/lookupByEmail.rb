require 'slack-ruby-client'
require "csv"

Slack.configure do |conf|
  conf.token = ""
end

client = Slack::Web::Client.new

rows = CSV.read("slack/tech_bowl/users (1).csv")
rows.each do |row|
  p row[14]
  row[7] = client.users_lookupByEmail(email: row[14]).user.id rescue nil
  sleep 1
end

users = rows.select do |row|
  row[7] != nil
end

CSV.open("slack/tech_bowl/users(registered_slack_id).csv", "w") do |csv|
  users.each do |user|
    csv << user
  end
end
