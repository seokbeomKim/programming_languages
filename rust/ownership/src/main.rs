// fn main() {
//     let s = "hello";
// }

fn main() {
    let mut s = String::from("hello");
    s.push_str(", world!");
    println!("{}", s);

    // rust 에서는 메모리 관리를 위한, GC가 없다. 대신 scope 에 따라 메모리가
    // 자동으로 해제된다.. function scope 인가? block scope인가?
    let x = 5;
    let y = x;

    // shallow copy - 이 부분은 일반적인 c/cpp와 같다. shallow copy가 아니었다?
    // rust에서는 shallow copy 대신 기존 s1을 해제하고 s2로 move 한다.
    let s1 = String::from("hello");
    let s2 = s1;
    // c/cpp와 같이 shallow copy를 인정하기보다 아예 s1 을 invalidate 시켜버리고
    // s2로 할당한다. 즉, shallow copy 가 발생하는 상황 자체를 인정하지 않고
    // rust 에서 말하는 ownership이 이걸까?
    // reference count 1 짜리가 ownership 아닌가?
    println!("{}, world!", s2);

    let len = calculate_length(&s2);
    println!("The length of {} is {}.", s2, len);
}

// 아래 함수에서 s: &String이 아닌 s: String으로 하게 되면 move 가 일어나게 되고
// calculate_length 함수가 끝나면 해당 메모리가 해제되어 버린다. 이처럼 함수의
// 인자로 reference를 가져오는 것을 borrowing이라고 한다.
fn calculate_length(s: &String) -> usize {
    s.len()
}