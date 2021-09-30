function doPost(r) {
  send_message();
}

function send_message(message) {
  let option = {
    method: "post",
    payload: buildMessage(),
    contentType: "application/x-www-form-urlencoded; charset=utf-8",
    muteHttpExceptions: true,
    Authorization: "Bearer",
  };

  let response = UrlFetchApp.fetch(
    "https://api.line.me/v2/bot/message/reply",
    option
  );
  Logger.log(response);
}

function buildMessage() {
    const bodyMessage = {
    "replyToken": ["replyToken"],
      "messages": [
        {
          "type": "text",
          "text": message
        }
      ]
    }

  return JSON.stringify(bodyMessage);
}

function message() {
  "こんにちは"
}