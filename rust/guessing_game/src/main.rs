// Using io library (which is known as std)
use std::cmp::Ordering;
use std::io;
use rand::Rng;

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..101);

    loop {
        println!("Please input your guess.");
        // 사용자 입력을 저장하기 위한 메모리 할당
        // 여기서 let과 mut은 어떤 의미인가?
        // let은 변수 할당 시 사용하는 키워드이다.
        // mut은 mutable을 의미하며, mut 키워드 없이 선언 시, immutable을 의미한다.
        let mut guess = String::new();
        // let test = 0;

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        // 이미 guess를 이전에 정의했는데도 불구하고 이전 변수를 shadow 할 수 있다.
        // let guess: u32 = guess.trim().parse().expect("Please type a number!");
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("You guessed: {}", guess);

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}
