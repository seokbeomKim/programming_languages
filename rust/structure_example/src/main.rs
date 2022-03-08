fn main() {
    main3();
}

// version 1
fn main1() {
    let width1 = 30;
    let height1 = 50;

    println!(
        "The area of the rectangle is {} square pixels.",
        area(width1, height1)
    );
}

fn area(width: u32, height: u32) -> u32 {
    width * height
}

// version 2
fn main2() {
    let rect1 = (30, 50);
    let x = 10;

    test_mov(x);
    println!("after test_mov() : {}", x);

    // 스택 변수들은 move와 관련이 없다. 즉, tuple 또한 object type이 아닌
    // primary type으로 분류 가능하다.

    println!(
        "The area of the rectangle is {} square pixels.",
        area2_1(rect1)
    );

    println!(
        "Still can print out dimension: {} {}",
        rect1.0, rect1.1
    )
}

fn test_mov(num: u32) {
    println!("number is {}", num);
}

fn area2(dimensions: &(u32, u32)) -> u32 {
    dimensions.0 * dimensions.1
}

// 왜 move가 안일어나지?
fn area2_1(dimensions: (u32, u32)) -> u32 {
    // println!("why? {} {}", dimensions.0 * dimensions.1);
    dimensions.0 * dimensions.1
}

// version 3: use
#[derive(Debug)]
// 구조체에 상속을 추가할 수 있는건가? 어떻게 가능한가?
struct Rectangle {
    width: u32,
    height: u32,
}

// vtable 스타일과.. python 스타일을 합친 느낌이다.
impl Rectangle {
    // 메서드에서도 레퍼런스로 정의해야 한다. move가 일어난다.
    fn area(&self) -> u32 {
        self.width * self.height
    }
}

// structure 는 인스턴스화 되므로 move 가 일어날 것이다 (예상)
fn main3() {
    let rect1 = Rectangle { width: 30, height: 50 };
    println!(
        "The area of the rectangle is {} square pixels.",
        area3(&rect1)
    );

    println!("{:?}", rect1);

    // 호출은 rect1.area() 형태로 하지만 사실 (&rect).area()의 의미와 같다.
    println!("{:?}", rect1.area());

    // 예상대로 instance화 된 rect1에 대해 reference로 넘기지 않으면 안된다.
    println!(
        "After invoke function, width: {}, height: {}",
        rect1.width, rect1.height
    );
}

fn area3(rect: &Rectangle) -> u32 {
    rect.width * rect.height
}