use scraper::{Html, Selector};

fn main() -> eyre::Result<()>{
    let body  = reqwest::blocking::get("https://blog.rust-lang.org/")?.text()?;

    let selector = Selector::parse("td.bn > a").unwrap();


    let document = Html::parse_document(&body);

    let elements = document.select(&selector);

    // 全記事名を出力
    elements.for_each(|e| println!("{}", e.text().next().unwrap()));

    Ok(())
}