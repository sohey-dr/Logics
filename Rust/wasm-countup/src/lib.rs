use js_sys::{Date};
use wasm_bindgen::prelude::*;
use wasm_bindgen::JsCast;
use web_sys::{Document, Element, HtmlElement, Window};

#[wasm_bindgen(start)]
pub fn run() -> Result<(), JsValue> {
    let window = web_sys::window().expect("should have a window in this context");
    let document = window.document().expect("window should have a document");

    setup_clock(&window, &document)?;

    document
        .get_element_by_id("script")
        .expect("should have #script on the page")
        .dyn_ref::<HtmlElement>()
        .expect("#script should be an `HtmlElement`")
        .style()
        .set_property("display", "block")?;

    Ok(())
}


fn setup_clock(window: &Window, document: &Document) -> Result<(), JsValue> {
    let current_time = document
        .get_element_by_id("current-time")
        .expect("should have #current-time on the page");
    update_time(&current_time);
    let a = Closure::wrap(Box::new(move || update_time(&current_time)) as Box<dyn Fn()>);
    window
        .set_interval_with_callback_and_timeout_and_arguments_0(a.as_ref().unchecked_ref(), 1000)?;
    fn update_time(current_time: &Element) {
        let date = Date::new_0();
        let mut seconds = date.get_seconds().to_string();
        if seconds.len() == 1 {
            seconds = format!("0{}", seconds);
        }

        current_time.set_inner_html(&format!(
            "{}:{}:{}",
            date.get_hours().to_string(),
            date.get_minutes().to_string(),
            seconds
        ));
    }
    a.forget();

    Ok(())
}
