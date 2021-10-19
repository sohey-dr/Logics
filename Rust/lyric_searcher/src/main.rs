use scraper::{Html, Selector};

fn main() -> eyre::Result<()>{
    let body = reqwest::blocking::get("https://www.google.com/search?q=%E5%90%9B%E3%81%8C%E5%A5%BD%E3%81%8D+%E6%AD%8C%E8%A9%9E")?.text()?;

    let selector = Selector::parse(".kCrYT > a").unwrap();

    let document = Html::parse_document(&body);

    for node in document.select(&selector) {
        println!("{:?}", node.inner_html());
        let href = node.value().attr("href").unwrap();

        if href.contains("uta-net") || href.contains("j-lyric.net") || href.contains("utamap") {
            println!("{:?}", href);
        }
    }

    Ok(())
}
