use std::{thread, time};

fn main() {
    println!("11111111");
    let ten_millis = time::Duration::from_millis(7000);
    thread::sleep(ten_millis);
    println!("Hello, world!");
}
