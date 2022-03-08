enum IpAddrKind {
    V4(String),
    V6(String),
}
struct IpAddr {
    kind: IpAddrKind,
    address: String,
}

enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

// enum 인데 인터페이스 정의가 가능한가?
impl Message {
    fn call(&self) {
        // method body would be ..
        println!("why?");
    }
}

fn route(ip_kind: IpAddrKind) {
    //
}

// Rust에서는 Null 을 허용하지 않는다. 대신 존재/비존재 를 나타내기 위해
// Option을 사용하며 아래와 같이 나타낼 수 있다. struct와 enum의 차이는 명시적인
// 필드를 사용하거나 match를 사용하는 차이점이다.
enum Option<T> {
    Some(T),
    None,
}

fn main() {
    let four = IpAddrKind::V4(String::from("Test"));
    let six = IpAddrKind::V6;

    let four = Message::Move{x:1, y:2};
    // println!("{}", four);
    four.call();

    let home = IpAddr {
        kind: IpAddrKind::V4(String::from("127.0.0.1")),
        address: String::from("127.0.0.1"),
    };

}
