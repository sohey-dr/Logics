(async() => {
  const HeadlessCrawler = require("./libs/HeadlessCrawler")
  const Options = require("./libs/Options")
  const options = new Options().run()
  const crawler = new HeadlessCrawler(options)
  await crawler.run()
})();