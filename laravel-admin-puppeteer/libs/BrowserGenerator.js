
const puppeteer = require('puppeteer');

class BrowserGenerator {
  constructor(options) {
    this.options = options
  }
  async generateBrowser() {
    const args = ['--no-sandbox', '--disable-gpu', '--single-process']
    if(this.options.proxy_url) args.push(`--proxy-server=${this.options.proxy_url}`)
    const browser = await puppeteer.launch({
      args: args,
      headless: !this.options.non_headless,
      slowMo: this.options.slow_mo
    });
    return browser;
  }

  async generatePage(browser) {
    const page = await browser.newPage();
    // const page = (await browser.pages())[0];
    // headlessモードとそうでない場合とでUAが違うみたいなので偽装が必要かもいれない
    await page.setUserAgent('Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36')
    if(this.options.proxy_url) {
      await page.authenticate({
        username: this.options.proxy_user,
        password: this.options.proxy_password
      })
    }
    return page;
  }
}

module.exports = BrowserGenerator