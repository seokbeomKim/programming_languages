// 아래에서 에러가 나는 이유는 이 함수에서 반환하는 값이 정확하게 어떤
// lifetime을 가지고 있지 알지 못하기 때문이다.

// 함수가 T 타입을 이용하여 generic type을 받도록 하듯이 아래와 같이 '(quote)를
// 이용하여 모든 lifetime 을 받을 수 있도록 정의할 수 있다. 아래의 'a는 x, y
// 모두 generic lifetime 동안 유지된다는 의미를 갖고 있다.
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn main() {

/*
    let r;

    {
        let x = 5;
        // x의 경우 lifetime이 inner block 까지만으므로 println이 호출되는
        // 시점까지 메모리가 유지되지 않는다. 이러한 것은 rust 컴파일러에서
        // borrow checker 를 가지고 있기 때문이다. borrow checker는 모둔 borrow
        // 들이 valid 한지를 체크한다.
        r = &x;
    }

    println!("{}", r);
*/

    let string1 = String::from("asdf");
    {
        let string2 = String::from("xyz");

        let result = longest(string1.as_str(), string2.as_str());
        println!("The longest string is {}", result);
    }
}
