use std::sync::mpsc::Receiver;

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }

    // #[test]
//     fn another() {
//         panic!("Make this test fail");
//     }
}

pub struct Rectangle {
    length: u32,
    width: u32,
}

impl Rectangle {
    pub fn can_hold(&self, other: &Rectangle) -> bool {
        self.length > other.length && self.width > other.width
    }
}

#[cfg(test)]
mod test {
    // 앞서 정의한 Rectangle을 사용하기 위해 아래와 같이 필요하다. visibility
    // 규칙에 따라 아래의 키워드 없이 곧바로 Rectangle 에 접근할 수 없기 때문
    use super::*;

    #[test]
    fn larget_can_hold_smaller() {
        let larger = Rectangle { length: 8, width: 7 };
        let smaller = Rectangle { length: 5, width: 1 };

        assert!(larger.can_hold(&smaller));

    }
}