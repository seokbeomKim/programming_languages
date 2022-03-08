struct Point<T> {
    x: T,
    y: T,
}

struct Point2<T, T2> {
    x: T,
    y: T2,
}

// 옵션의 경우 아래와 같이 정의된다.
enum Option<T> {
    Some(T),
    None,
}

impl<T> Point<T> {
    fn x(&self) -> &T {
        &self.x
    }
}


fn main() {
    // 아래와 같이 i32, float 타입 동시에 가질 수 없다.
    // 그렇다면, T, T2 등 여러 개의 타입으로 정의할 수 있을까?
    // let wont_work = Point {x:5,y:4.0};

    let p = Point {x: 5, y: 10};
    // 동작한다!!
    let does_work = Point2 {x:5, y:4.0};

    println!("p.x = {}", p.x());
}
