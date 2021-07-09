const fs = require('fs')
const BrowserGenerator = require('./BrowserGenerator')

module.exports = class HeadlessCrawler {
  constructor(options) {
    this.options = options
    //
    this.browser = null;
    this.page = null;
    //
    this._mkdirP(this._outputDir());
  }

  async run() {
    try {
      await this._generateBrowserAndPage()
      await this.page.goto(this._url())
      const html = await this.page.$eval('html', item => {
        return item.innerHTML
      })

      this._saveHtml(html)
    } catch(e) {
      console.log(e)
      this._saveError(e)
    }
    await this.browser.close()
  }

  _saveHtml(html) {
    fs.writeFileSync(this._filepath(), html);
  }

  _saveError(e) {
    fs.writeFileSync(this._errorFilepath(), e);
  }

  _mkdirP(dir) {
    if (!fs.existsSync(dir)) {
      fs.mkdirSync(dir);
    }
  }

  _filepath() {
    const filepath = `${this._outputDir()}/crawl.html`
    return filepath
  }

  _errorFilepath() {
    const filepath = `${this._outputDir()}/ERROR`
    return filepath
  }

  async _generateBrowserAndPage() {
    const browserGenerator = new BrowserGenerator(this.options)
    this.browser = await browserGenerator.generateBrowser()
    this.page = await browserGenerator.generatePage(this.browser)
  }

  _url() {
    return this.options.url
  }

  _outputDir() {
    return this.options.output
  }
}