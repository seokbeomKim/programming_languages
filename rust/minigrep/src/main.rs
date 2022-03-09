// bring std::env module into scope with a use statment so we can use its args
// function.
use std::env;
use std::process;

use minigrep::Config;

fn main() {
    // 만약 인자안에 invalid Unicode가 섞여 있는 경우 아래에서 std::env::args 는
    // 패닉을 발생시킨다.
    let args: Vec<String> = env::args().collect();
    // 첫 번째 인자는 ["target/debug/minigrep", "searchstring", "테스트"] 와
    // 같이 실행되는 바이너리 경로가 된다.
    // println!("{:?}", args);

    // unwrap_or_else 를 이용하여 Result<T, E> 로 나타낼 수 있다.
    let config = Config::new(&args).unwrap_or_else(|err| {
        println!("Problem parsing arguments: {}", err);
        process::exit(1);
    });

    if let Err(e) = minigrep::run(config) {
        println!("Application error: {}", e);

        process::exit(1);
    }
}

#[allow(dead_code)]
fn parse_config(args: &[String]) -> Config {
    // 아래와 같이 clone을 하게 되는 경우 lifetime을 별도로 관리할 필요가
    // 없어지기 때문에 편리하지만 추가적인 메모리가 필요하게 된다. 이렇게
    // clone하는 것은 런타임에서의 비용이 발생하므로 이를 위한 다른 대안이
    // 필요하다. (추후 설명)
    let query = args[1].clone();
    let filename = args[2].clone();

    Config { query, filename }
}
