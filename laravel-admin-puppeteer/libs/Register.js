const puppeteer = require('puppeteer');

class Register {
  constructor() {}

  async run() {
    const browser = await puppeteer.launch({
      headless: false,
      slowMo: 500,
    });

    const page = await browser.newPage();
    await page.goto("https://www.google.com/");

    await browser.close();
  }
}

module.exports = Register