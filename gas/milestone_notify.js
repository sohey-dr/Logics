function run() {
  // notify();
  const milestones = Object.values(
    getMilestones(["laravel-backend", "website"])
  );
  let spreadSheet = SpreadsheetApp.openById(
    ""
  );
  let targetSheet = spreadSheet.getSheetByName(getYearMonth());
  if (targetSheet === null) {
    targetSheet = createSheet();
  }

  Logger.log(milestones);
  targetSheet.getRange("B2").setValues("test");
}

function createSheet() {
  const newCopySheet = spreadSheet
    .getSheetByName("元シート")
    .copyTo(spreadSheet);
  newCopySheet.setName(getYearMonth());

  return newCopySheet;
}

function getYearMonth() {
  let date = new Date();
  return Utilities.formatDate(date, "Asia/Tokyo", "yyyy/MM");
}

function setTrigger() {
  // 時間指定の定期実行トリガーだと1時間の幅ができてしまうため本メソッドを実行
  const next = new Date();
  next.setDate(next.getDate() + 7);
  next.setHours(9);
  next.setMinutes(00);
  next.setSeconds(0);
  ScriptApp.newTrigger("notify").timeBased().at(next).create();
}

// Slack系の処理
function notify() {
  const milestones = getMilestones(["laravel-backend", "website"]);
  const messages = buildMessage(milestones);

  Logger.log(messages);
  send_message(messages);
  // setTrigger();
}

function slackWebhookUrl() {
  return PropertiesService.getScriptProperties().getProperty(
    "SLACK_WEBHOOK_URL"
  );
}

function send_message(messages) {
  let url = slackWebhookUrl();
  let payload = { blocks: messages };
  let option = {
    method: "post",
    payload: JSON.stringify(payload),
    contentType: "application/x-www-form-urlencoded; charset=utf-8",
    muteHttpExceptions: true,
  };
  let response = UrlFetchApp.fetch(url, option);
  Logger.log(response);
}

function buildMessage(milestones) {
  let content = [
    {
      type: "section",
      text: { type: "mrkdwn", text: "*今週の開発の進捗です*\n<!here>" },
    },
  ];
  let messages = [];
  for (const milestone of milestones) {
    let progressPercentage = Math.round(milestone["progressPercentage"]);
    let message = {
      type: "mrkdwn",
      text: `*<${milestone["url"]}|${milestone["title"]}>*\n*完了:* ${milestone["closed"]}個\n*未完了:* ${milestone["open"]}個\n達成率:${progressPercentage}%\n_`,
    };
    messages.push(message);
  }
  const bodyMessage = {
    type: "section",
    fields: messages,
  };
  content.push(bodyMessage);

  return JSON.stringify(content);
}

// GitHub系の処理
function githubAccessToken() {
  return PropertiesService.getScriptProperties().getProperty(
    "GITHUB_ACCESS_TOKEN"
  );
}

function getMilestones(repositories) {
  let milestonesList = [];
  for (const repository of repositories) {
    const response = fetch_pullreq_data_by_graphql(repository);
    const json = response.getContentText();
    const parsedJson = JSON.parse(json);
    let milestones = parsedJson.data.repository.milestones.nodes;

    for (const milestone of milestones) {
      issuesStatus = issuesCount(milestone.issues.nodes);
      milestones[milestones.indexOf(milestone)] = Object.assign(
        milestones[milestones.indexOf(milestone)],
        issuesStatus
      );
      delete milestones[milestones.indexOf(milestone)].issues;
    }

    Array.prototype.push.apply(milestonesList, milestones);
  }

  return milestonesList;
}

function issuesCount(issues) {
  open = 0;
  closed = 0;
  for (const issue of issues) {
    switch (issue.state) {
      case "OPEN":
        open++;
        break;
      case "CLOSED":
        closed++;
        break;
      default:
        console.log("Erorr");
    }
  }

  return { open: open, closed: closed };
}

function fetch_pullreq_data_by_graphql(repository) {
  const graphql_query =
    '{\
    repository(owner: "TechBowl-japan", name: "' +
    repository +
    '") {\
      milestones(states: [OPEN], last: 20) {\
        nodes {\
          title\
          progressPercentage\
          url\
          issues(last: 60) {\
            nodes {\
              state\
            }\
          }\
        }\
      }\
    }\
  }';

  const option = buildRequestOption(graphql_query);
  return UrlFetchApp.fetch("https://api.github.com/graphql", option);
}

function buildRequestOption(graphql) {
  return {
    method: "post",
    contentType: "application/json",
    headers: {
      Authorization: "bearer " + githubAccessToken(),
    },
    payload: JSON.stringify({ query: graphql }),
  };
}
