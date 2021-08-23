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
    const date = new Date();
    const today = Utilities.formatDate(date, "Asia/Tokyo", "MM/dd");

    let targetCell = null;
    const lastRow = sheet.getRange(sheet.getLastRow(), 1);
    const lastRowDate = Utilities.formatDate(
      new Date(lastRow.getValue()),
      "Asia/Tokyo",
      "MM/dd"
    );

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
