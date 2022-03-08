use std::{fs::File, io::{ErrorKind, Read}};

fn read_username_from_file() -> Result<String, io::Error> {
    let f = File::open("hello.txt");

    let mut f = match f {
        Ok(file) => file,
        Err(e) => return Err(e),
    };

    let mut s = String::new();
    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}

fn read_username_from_file() -> Result<String, io::Error> {
    let mut f = File::open("hello.txt");
    let mut s = String::new();
    // 아래와 같이 물음표? 를 사용할 경우 에러가 발생하면 함수는 Err를 리턴한다.
    f.read_to_string(&mut s)?;
    Ok(s)
}

fn main() {
    let f = File::open("hello.txt");

    // let f = match f {
    //     Ok(file) => file,
    //     Err(error) => match error.kind(){
    //         ErrorKind::NotFound => match File::create("hello.txt") {
    //             Ok(fc) => fc,
    //             Err(e) => panic!("Problem creating the file: {:?}", e),
    //         },
    //         other_error => panic!("Problem opening the file: {:?}", other_error),
    //     },
    // };
    // let f = File::open("hello.txt").unwrap();
    let f = File::open("hello.txt").expect("Failed to open hello.txt");
}