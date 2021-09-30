const CHANNEL_ACCESS_TOKEN = "";

const MESSAGE = "これは届いていますか？";

function doPost(r) {
  const params = JSON.parse(r.postData.getDataAsString()).events[0];
  send_message(params);
}

function send_message(params) {
  var headers = {
    Authorization: "Bearer " + CHANNEL_ACCESS_TOKEN,
  };

  let option = {
    method: "post",
    headers: headers,
    payload: buildMessage(params),
    contentType: "application/json",
    muteHttpExceptions: true,
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
        text: MESSAGE,
      },
    ],
  };

  return JSON.stringify(bodyMessage);
}
