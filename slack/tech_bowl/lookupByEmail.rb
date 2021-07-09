require 'slack-ruby-client'

Slack.configure do |conf|
  conf.token = ""
end

client = Slack::Web::Client.new
user = client.users_lookupByEmail(email: "heyso1209@gmail.com").user
user.id