use js_sys::{Date};
use wasm_bindgen::prelude::*;
use wasm_bindgen::JsCast;
use web_sys::{Document, Element, Window};

#[wasm_bindgen(start)]
pub fn run() -> Result<(), JsValue> {
    let window = web_sys::window().expect("should have a window in this context");
    let document = window.document().expect("window should have a document");

    setup_clock(&window, &document)?;

    Ok(())
}


fn setup_clock(window: &Window, document: &Document) -> Result<(), JsValue> {
    let year = document
        .get_element_by_id("year")
        .expect("should have #time on the page");
    let days = document
        .get_element_by_id("days")
        .expect("should have #time on the page");
    let time = document
        .get_element_by_id("time")
        .expect("should have #time on the page");
    // 読み込み時に時間を取得
    update_time(&year, &days, &time);

    // 毎秒時計を更新
    let a = Closure::wrap(Box::new(move || update_time(&year, &days, &time)) as Box<dyn Fn()>);
    window
        .set_interval_with_callback_and_timeout_and_arguments_0(a.as_ref().unchecked_ref(), 1000)?;
    fn update_time(year: &Element, days: &Element, time: &Element) {
        let date = Date::new_0();

        year.set_inner_html(&date.get_full_year().to_string());

        days.set_inner_html(&format!(
            "{}/{}",
            date.get_month(),
            date.get_day()));

        let mut seconds = date.get_seconds().to_string();
        if seconds.len() == 1 {
            seconds = format!("0{}", seconds);
        }
        time.set_inner_html(&format!(
            "{}:{}:{}",
            date.get_hours().to_string(),
            date.get_minutes().to_string(),
            seconds
        ));
    }
    a.forget();

    Ok(())
}
