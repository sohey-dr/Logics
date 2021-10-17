use scraper::{Html, Selector};

fn main() {
    let body  = reqwest::blocking::get("https://blog.rust-lang.org/")?.text()?;

    // HTMLをパース
    let document = scraper::Html::parse_document(&body);

    let fragment = Html::parse_fragment(document);
    let selector = Selector::parse("li").unwrap();

    for element in fragment.select(&selector) {
        assert_eq!("li", element.value().name());
    }
}