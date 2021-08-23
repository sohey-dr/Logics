require 'slack-ruby-client'

Slack.configure do |conf|
  conf.token = ""
end

client = Slack::Web::Client.new

client.users_lookupByEmail(email: "")
