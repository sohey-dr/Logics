function createLog() {
  const milestones = Object.values(
    getMilestones(["laravel-backend", "website"])
  );
  let spreadSheet = SpreadsheetApp.openById(
    "1dyvAONQkyZQ5ovj07IdAdHVcZKREV-QqAlh4bYRfV0w"
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
