const { Scanner, Utils } = require("spherov2.js");

const makeItBlink = async () => {
  const sphero = await Scanner.findSpheroMini();

  if (!sphero) return console.log("sphero mini not available!");

  while (true) {
    await sphero.setMainLedColor(255, 0, 0);
    await Utils.wait(200);
    await sphero.setMainLedColor(255, 255, 0);
    await Utils.wait(200);
    await sphero.setMainLedColor(0, 255, 0);
    await Utils.wait(200);
    await sphero.setMainLedColor(0, 255, 255);
    await Utils.wait(200);
    await sphero.setMainLedColor(0, 0, 255);
    await Utils.wait(200);
    await sphero.setMainLedColor(255, 0, 255);
    await Utils.wait(200);
  }
};

makeItBlink();
