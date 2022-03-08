fn main() {
    let reference_to_nothing = dangle2();
    let fw = first_word_using_slice(&reference_to_nothing);
    println!("{}", fw);

    let s = String::from("hello world");
    let hello = &s[0..5];
    let world = &s[6..11];

    println!("{}, {}", hello, world);
}

/*
// lifetime 에러 발생?
fn dangle() -> &String {    // dangle returns a reference to a String
    let s = String::from("hello");

    &s
}   // 여기서 scope를 벗어나는 순간 할당한 s가 메모리에서 해제된다.
*/

fn dangle2() -> String {
    // 이 경우에는 ownership이 함께 전달되기 때문에 (이전에 인자 값으로 넘겼을
    // 때와 마찬가지로) 아래와 같이 전달해줘야 한다.
    let s = String::from("hello, world");

    s
}
fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }

    s.len()
}

fn first_word_using_slice(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[..i];
        }
    }

    &s[..]
}

