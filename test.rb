require 'csv'

rows = CSV.read("slack/tech_bowl/users(registered_slack_id).csv")

users = rows.select do |row|
  row[7] != nil
end

CSV.open("slack/tech_bowl/users(registered_slack_id).csv", "w") do |csv|
  users.each do |user|
    csv << user
  end
end