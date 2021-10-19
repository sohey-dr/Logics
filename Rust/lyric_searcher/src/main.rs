use scraper::{Html, Selector};

fn main() -> eyre::Result<()>{
    let body = reqwest::blocking::get("https://www.google.com/search?q=%E5%90%9B%E3%81%8C%E5%A5%BD%E3%81%8D+%E6%AD%8C%E8%A9%9E")?.text()?;
    let document = Html::parse_document(&body);
    let selector = Selector::parse(".kCrYT > a").unwrap();

    for node in document.select(&selector) {
        let href = node.value().attr("href").unwrap();

        let last_index = href.find("&sa=").unwrap();
        let url = &href[7..last_index];

        if url.contains("uta-net") || url.contains("j-lyric.net") || url.contains("utamap") {
            println!("{:?}", node.text().collect::<Vec<_>>()[0]);
            println!("{:?}", url);
        }
    }

    Ok(())
}
