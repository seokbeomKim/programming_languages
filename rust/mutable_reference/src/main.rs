fn main() {
    let s = String::from("hello");

    // 이 경우에는 shallow copy가 아니라 아예 move가 일어나버려 기존 s는 사용할
    // 수 없게 된다.
    // let s2 = s;
    // change(s);

    // 흥미로운건 immutable 타입의 값을 mutable 로 가져올 수 있다.
    // mutable_reference 경고가 발생하지만 말이다.
    let mut s3 = s;
    change3(&mut s3);

    // 아래 또한 제대로 동작하지 않는다. change 함수의 인자로 이미 move가
    // 일어났기 때문에 함수 동작 이후에 기존 s가 가지고 있던 값은 메모리
    // 해제된다.
    println!("{}", s3);

    // 아래는 s3 에 의해 이미 move가 한 번 일어났으므로 사용이 불가능하다.
    // println!("{}", s);
}

// main에서 함수 호출 시 인자로 넘긴 값이 move 일어나는 버전
fn change(some_string: String) {
    // some_string.push_str(", world");
    println!("{}", some_string);
}

// main에서 함수 호출 이후에도 인자로 인해 move가 일어나지 않는 버전
// reference를 통해 인자로 넣어줘야 한다. 즉, c로 비유하면 포인터 사용을 강제하고 있다.
// 하지만 이 경우 some_string 인자로 넘어온 문자열을 수정할 수 없다..!?
fn change2(some_string: &String) {
    // 아래 함수는 동작하지 않는다.
    // some_string.push_str(", world");
    println!("{}", some_string);
}

fn change3(some_string: &mut String) {
    some_string.push_str(", world");
    println!("{}", some_string);
}