// const가 있는데 immutable 변수가 사용되는 이유는 무엇인가? 컴파일타임
// constant, 런타임 constant 의 차이점을 두기 위해 설계된 키워드들이다. 굳이
// constant 를 컴파일 타임, 런타임으로 구분할 필요가 있었을까 하는 생각이
// 들지만, 어떤 함수 내에서 문맥의 의미를 간결하게 하기 위해 가끔씩은 유용하게
// 사용할 수 있을 것 같다.
const MAX_POINT: u32 = 100_000;
fn main() {
    // chapter 3. Common Programming Concepts에 나오는 예제들을 아래부터 작성한다.
    let mut x = 5;
    x = MAX_POINT;
    // 왜 mutual 을 명시적으로 한 걸까? const 대용인가?
    println!("The value of x is {}", x);

    test();
}

fn test() {
    // let을 이용하여 변수를 매번 선언하는게 아니라, 한 번 선언되면 스택에
    // 메모리가 할당되는 방식인가? 출력 결과는 아래와 같다. 즉, shadowing 자체는
    // 새로운 변수 할당과 같고 syntax 상으로 완전히 가리게 된다.
    //
    // address of x: 0x7ffd3d4ee86c
    // address of x: 0x7ffd3d4ee8c4
    // address of x: 0x7ffd3d4ee91c
    let x = 5;
    println!("address of x: {:p}", &x);
    let x = x + 1;
    println!("address of x: {:p}", &x);
    let x = x * 2;
    println!("address of x: {:p}", &x);
    println!("The value of x is : {}", x);

    let spaces = "    ";
    // let spaces = spaces.len();
    // type이 string과 int로 서로 다르므로 할당 불가능
    // spaces = spaces.len();
    let tup: (i32, f64, u8) = (500, 6.4, 1);
    let (x, y, z)= tup;
    println!("The value of y is: {}", y);

    let x = tup;
    // 튜플에서 인덱스를 통해 변수를 얻어올 수 있다..
    let five_hundred = x.0;
    let six_point_four = x.1;
    let one = x.2;

    println!("{} {} {}", five_hundred, six_point_four, one);
}
