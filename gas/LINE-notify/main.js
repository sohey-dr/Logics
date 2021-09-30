function doPost(r) {
  const params = JSON.parse(r.postData.getDataAsString());
  send_message(params);
}

function send_message(params) {
  let option = {
    method: "post",
    payload: buildMessage(params),
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

function buildMessage(params) {
  const bodyMessage = {
    replyToken: params.replyToken,
    messages: [
      {
        type: "text",
        text: message,
      },
    ],
  };

  return JSON.stringify(bodyMessage);
}

function message() {
  "こんにちは"
}