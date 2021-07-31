// Slack系の処理
function notifySlack() {
  const milestones = getMilestones(["laravel-backend", "website"]);
  const messages = buildMessage(milestones);

  send_message(messages);
}

function slackWebhookUrl() {
  return PropertiesService.getScriptProperties().getProperty("SLACK_WEBHOOK_URL");
}

function send_message(messages) {
  let url = slackWebhookUrl();
  let payload = {"blocks": messages }
  let option = {
    'method': 'post',
    'payload': JSON.stringify(payload),
    'contentType': 'application/x-www-form-urlencoded; charset=utf-8',
    'muteHttpExceptions': true
  };
  let response = UrlFetchApp.fetch(url, option);
  Logger.log(response);
}

function buildMessage(milestones) {
  let content = [{"type": "section","text": {"type": "mrkdwn","text": "*今週の開発の進捗です*\n<!here>"}}];
  let messages = [];
  for (const milestone of milestones) {
    let progressPercentage = Math.round(milestone['progressPercentage'])
    let message = {
          "type": "mrkdwn",
          "text": `*<${milestone['url']}|${milestone['title']}>*\n*完了:* ${milestone['closed']}個\n*未完了:* ${milestone['open']}個\n*達成率:* ${progressPercentage}%\n_`
        }
    messages.push(message)
  }
  const bodyMessage = {
    "type": "section",
    "fields": messages
  }
  content.push(bodyMessage);

  return JSON.stringify(content);
}