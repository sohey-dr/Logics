function doPost(r) {
  const spreadSheet = SpreadsheetApp.openById(
    "1P_X5ghsn2VBKgkUItgYNZkGc3VSz5vnOcgKgTMXhPow"
  );
  const sheet = spreadSheet.getSheetByName("シート1");

  const params = JSON.parse(r.postData.getDataAsString());

  if (params.type === "url_verification") {
    // ここはSlackのVerify時用
    return ContentService.createTextOutput(params.challenge);
  } else if (params.event.type === "team_join") {
    let targetCell = null;
    const lastRow = sheet.getRange(sheet.getLastRow(), 1);
    const date = new Date();
    const today = Utilities.formatDate(date, "Asia/Tokyo", "MM/dd");
    const lastRowDate = Utilities.formatDate(
      new Date(lastRow.getValue()),
      "Asia/Tokyo",
      "MM/dd"
    );

    // もしその日のものがなかったら新しく行を作る
    if (today === lastRowDate) {
      targetCell = lastRow.offset(0, 1);
    } else {
      lastRow.offset(1, 0).setValue(today);
      targetCell = lastRow.offset(1, 1);
    }

    targetCell.setValue(targetCell.getValue() + 1);

    return null;
  }
}

function run() {
  const spreadSheet = SpreadsheetApp.openById(
    "1P_X5ghsn2VBKgkUItgYNZkGc3VSz5vnOcgKgTMXhPow"
  );
  const sheet = spreadSheet.getSheetByName("シート1");
  const date = new Date();
  const today = Utilities.formatDate(date, "Asia/Tokyo", "MM/dd");
  const lastRowDate = Utilities.formatDate(
    new Date(sheet.getRange(sheet.getLastRow(), 1).getValue()),
    "Asia/Tokyo",
    "MM/dd"
  );

  const todayNumber = sheet.getRange(sheet.getLastRow(), 2).getValue();
  const existJoinUser = today === lastRowDate;
  const message = buildMessage(todayNumber, existJoinUser);
  send_message(message);
  setTrigger();
}

function slackWebhookUrl() {
  return PropertiesService.getScriptProperties().getProperty(
    "SLACK_WEBHOOK_URL"
  );
}

function send_message(message) {
  let url = slackWebhookUrl();
  let payload = { blocks: message };
  let option = {
    method: "post",
    payload: JSON.stringify(payload),
    contentType: "application/x-www-form-urlencoded; charset=utf-8",
    muteHttpExceptions: true,
  };
  let response = UrlFetchApp.fetch(url, option);
  Logger.log(response);
}

function buildMessage(todayNumber, existJoinUser) {
  let bodyMessage = null;
  if (existJoinUser) {
    bodyMessage = [
      {
        type: "section",
        text: {
          type: "mrkdwn",
          text: "今日のRailwayワークスペース参加者は" + todayNumber + "人です",
        },
      },
    ];
  } else {
    bodyMessage = [
      {
        type: "section",
        text: {
          type: "mrkdwn",
          text: "今日のRailwayワークスペース参加者は0人です",
        },
      },
    ];
  }

  return JSON.stringify(bodyMessage);
}

function slackWebhookUrl() {
  return PropertiesService.getScriptProperties().getProperty(
    "SLACK_WEBHOOK_URL"
  );
}

function setTrigger() {
  // 時間指定の定期実行トリガーだと1時間の幅ができてしまうため本メソッドを実行
  const next = new Date();
  next.setDate(next.getDate() + 1);
  next.setHours(23);
  next.setMinutes(58);
  next.setSeconds(0);
  ScriptApp.newTrigger("run").timeBased().at(next).create();
}
