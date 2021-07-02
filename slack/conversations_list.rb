require 'slack-ruby-client'

Slack.configure do |conf|
  conf.token = ""
end

client = Slack::Web::Client.new
p client.conversations_list(limit: 1000).channels.size