require("dotenv").config();
const puppeteer = require("puppeteer");

class Register {
  constructor(user) {
    // userは配列
    this.userId = user[0];
    this.userPrivateRoomId = user[6];
    this.userSlackId = user[7];
  }

  async run() {
    const browser = await puppeteer.launch({
      headless: false,
      slowMo: 50,
    });

    const page = await browser.newPage();
    await page.goto(
      process.env.TECHBOWL_BASEURL +
        "/admin" +
        "/users/" +
        this.userId +
        "/edit"
    );

    await this._login(page);
    await page.waitForTimeout(1000);

    console.log("該当ユーザーID" + this.userId + ":" + this.userSlackId)

    await this._registerForSlack(page);

    await browser.close();
  }

  async _login(page) {
    await page.type('input[name="username"]', process.env.USER_NAME);
    await page.type('input[name="password"]', process.env.USER_PASS);
    await page.click('button[type="submit"]');
  }

  async _registerForSlack(page) {
    await page.type('input[name="slack_id"]', this.userSlackId);
    await page.type('input[name="messenger_tool_private_room_id"]', this.userPrivateRoomId);
    await page.click('button[type="submit"]');
  }
}

module.exports = Register;
