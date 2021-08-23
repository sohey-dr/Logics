function doPost(r) {
  const spreadSheet = SpreadsheetApp.openById(
    "1P_X5ghsn2VBKgkUItgYNZkGc3VSz5vnOcgKgTMXhPow"
  );
  const sheet = spreadSheet.getSheetByName("シート1");

  const params = JSON.parse(r.postData.getDataAsString());
  const eventType = params.event.type;

  sheet.getRange("A1").setValue(eventType);

  // return {"data": eventType}
}
