require 'slack-ruby-client'
require "csv"

Slack.configure do |conf|
  conf.token = ""
end

client = Slack::Web::Client.new

rows = CSV.read("slack/tech_bowl/users (1).csv")
rows.each do |row|
  row[7] = client.users_lookupByEmail(email: row[14]).user.id
end

CSV.open("slack/tech_bowl/users(registered_slack_id)", "w") do |csv|
  rows.each do |row|
    csv << row
  end
end
