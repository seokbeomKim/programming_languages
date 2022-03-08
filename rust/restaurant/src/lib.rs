#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}

// 모듈 정의를 위해서는 아래와 같이 `mod` 키워드를 이용한다.

mod front_of_house {
    pub mod hosting {
        pub fn add_to_waitlist() {}
        fn seat_at_table() {}
    }

    mod serving {
        fn take_order() {}
        fn serve_order() {}
        fn take_payment() {}
    }
}

use self::front_of_house::hosting;

pub fn eat_at_restaurant() {
    // crate::front_of_house::hosting::add_to_waitlist();
    hosting::add_to_waitlist();
}