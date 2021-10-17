use scraper::{Html, Selector};

fn main() -> eyre::Result<()>{
    let body  = reqwest::blocking::get("https://blog.rust-lang.org/")?.text()?;
    println!("{} を取得", body);

    let fragment = Html::parse_fragment(&body);
    let selector = Selector::parse("li").unwrap();

    for element in fragment.select(&selector) {
        assert_eq!("li", element.value().name());
    }
    Ok(())
}