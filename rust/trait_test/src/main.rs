
// Trait과 인터페이스의 차이는 무엇인가? 동일한 개념인가?
// pub trait Summary {
//     fn summarize(&self) -> String;
// }

// 아래와 같이 method의 기본 behavior 를 정의할 수 있다. 이 경우 overriding 하여
// 새로 구현할 수도 있다.

pub trait Summary {
    fn summarize(&self) -> String {
        String::from("(read more..)")
    }
}

pub struct NewsArticle {
    pub headline: String,
    pub location: String,
    pub author: String,
    pub content: String,
}

impl Summary for NewsArticle {
    fn summarize(&self) -> String {
        format!("{}, by {} ({})", self.headline, self.author, self.location)
    }
}

pub struct Tweet {
    pub username: String,
    pub content: String,
    pub reply: bool,
    pub retweet: bool,
}

// 인터페이스 구현
impl Summary for Tweet {
    // 아래의 summarize를 별도로 정의하지 않으면 trait 내에 정의된 기본 trait의
    // 메서드가 호출된다.
//     fn summarize(&self) -> String {
//         format!("{}: {}", self.username, self.content)
//     }
}

// 메서드 구현
impl Tweet {
    fn test(&self, t: i32) {
        println!("This is test method: {}", t)
    }
}

// 인터페이스를 인자로 사용함으로서 유연성을 가질 수 있다.
pub fn notify(item: impl Summary) {
    println!("Breaking news! {}", item.summarize())
}

fn largest<T: PartialOrd + Copy>(list: &[T]) -> T {
    let mut largest = list[0];

    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }

    largest
}
fn main() {
    println!("Hello, world!");
    let t = Tweet {
        username: String::from("sukbeom"),
        content: String::from("Hi!"),
        reply: true,
        retweet: true,
    };
    t.test(10);

    // println!("{}", t.summarize())

    let number_list = vec![34, 50, 25, 100, 65];
    let result = largest(&number_list);
    println!("The largest: {}", result)
}
