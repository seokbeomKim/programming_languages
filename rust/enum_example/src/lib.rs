// Rust has a number of features that allow you to manage the code's
// organization, including details are exposed, which details are prive, ...
// and what names are in each scope in the program.
// 즉, 모듈이라는 시스템(?)을 통해서 프로그램의 상세 내용을 정의할 수 있다. 어떻게?

// [관련 용어]
// crates: A tree of modules that produces a library or executable
// packages: 패키지 build, test, share crates
// modules and use: 각 경로에 있는 구성들을 control 할 수 있는 개념
// path: item, struct, function, module 등의 네이밍 방법

mod front_of_house {
        mod hosting {
                fn add_to_waitlist() {}
                fn seat_at_table() {}
        }

        mod serving {
                fn take_order() {}
                fn serve_order() {}
                fn take_payment() {}
        }
}

pub fn eat_at_restaurant() {
        crate::front_of_house::hosting::add_to_waitlist();

        front_of_house::hosting::add_to_waitlist();
}