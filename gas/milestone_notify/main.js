function run() {
  notify();

  //TODO:createLogでスプレットシートにログを残しておく
  setTrigger();
}

function setTrigger() {
  // 時間指定の定期実行トリガーだと1時間の幅ができてしまうため本メソッドを実行
  const next = new Date();
  next.setDate(next.getDate() + 7);
  next.setHours(9);
  next.setMinutes(00);
  next.setSeconds(0);
  ScriptApp.newTrigger("run").timeBased().at(next).create();
}
