function doPost(r) {
  const spreadSheet = SpreadsheetApp.openById(
    "1P_X5ghsn2VBKgkUItgYNZkGc3VSz5vnOcgKgTMXhPow"
  );
  const sheet = spreadSheet.getSheetByName("シート1");

  const params = JSON.parse(r.postData.contents);

  if (params.type == "url_verification") {
    return ContentService.createTextOutput(
      JSON.stringify(contents.challenge)
    ).setMimeType(ContentService.MimeType.TEXT);
  } else if (params.event.type !== "team_join") {
    return JSON.stringify({ message: "Not Team Join" });
  }

  const targetCell = sheet.getRange("B1");
  targetCell.setValue(Number(targetCell.getValue()) + 1);

  return;
}
