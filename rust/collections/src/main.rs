fn main() {
    // 벡터의 경우에는 원소들의 타입에 대해 반드시 명시해줘야 한다.
    let v: Vec<i32> = Vec::new();

    // 또는 아래와 같이 사용할 수 있다.
    let v = vec![1,2,3];

    let mut v = Vec::new();
    v.push(5);
    v.push(6);
    v.push(7);
    v.push(8);

    let third = &v[2];
    // let third: &i32 = &v[9];
    // println!("The third element is {}", third);

    // 아래는 좀 더 안전하게 벡터에 접근하는 방법이다.
    match v.get(9) {
        Some(third) => println!("The third element is {}", third),
        None => println!("There is no third element.")
    }

    let mut v = vec![1,2,3,4,5];
    let mut first = &v[0];

    // 아래 라인에서 에러가 나는데 (move 관련) 이는 vector의 동작 방식 때문에
    // 발생한다. 벡터에 새로운 항목을 추가하면 새로운 메모리를 할당하고 기존
    // 원소들을 새로운 메모리 공간에 copy 해야하는 일이 발생한다. 때문에 push
    // 하는 순간 기존 내용은 모두 해제되므로 이러한 상황이 발생하지 않도록
    // borrowing rule 이 적용된다.
    // v.push(6);

    // 아래 라인을 사용하지 않는 경우 기존 first 가지고 있는 레퍼런스는 move가
    // 일어나므로 무효화된다. let first = &v[0];
    println!("The first element is: {}", first);
}
