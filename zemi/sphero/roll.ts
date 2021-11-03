const { Scanner, Utils } = require("spherov2.js");

const makeItRoll = async () => {
  const sphero = await Scanner.findSpheroMini();

  if (!sphero) return console.log("sphero mini not available!");

  const speed = 100;
  const headingInDegrees = 0;
  const timeToRollInMilliseconds = 2000;

  await sphero.rollTime(
    speed,
    headingInDegrees,
    timeToRollInMilliseconds,
    [],
  );
};

makeItRoll();
