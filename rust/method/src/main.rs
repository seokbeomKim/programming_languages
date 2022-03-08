use std::sync::mpsc::Receiver;

struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {

    // Associated Function 이라고 하는데 constructor이다. 통상 이러한
    // constructor 부분과 method 부분을 구분하기 위해 impl block을 여러개
    // 허용한다.
    fn square(size: u32) -> Rectangle {
        Rectangle { width: size, height: size }
    }
}

impl Rectangle {
    fn can_hold(&self, rect: &Rectangle) -> bool {
        self.width * self.height > rect.width * rect.height
    }
}

fn main() {
    let rect1 = Rectangle {width: 30, height: 50};
    let rect2 = Rectangle {width: 10, height: 40};
    let rect3 = Rectangle {width: 60, height: 45};

    println!("Can rect1 hold rect2? {}", rect1.can_hold(&rect2));
    println!("Can rect1 hold rect3? {}", rect1.can_hold(&rect3));
}
