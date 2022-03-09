use std::error::Error;
use std::fs;

pub struct Config {
    pub query: String,
    pub filename: String,
}

// 아래의 parse_config 대신 객체화시켜서 아래와 같이 constructor 형태로 만들 수
// 있다.
impl Config {
    pub fn new(args: &[String]) -> Result<Config, &'static str> {
        if args.len() < 3 {
            // 이와 같이 panic 을 발생시켜버리면 사용자에게 노출하고 싶지 않은
            // 정보를 가릴 수 있고 대신, debug point를 얻어낼 수 있기에
            // 유용하다.
            // panic!("not enough arguments");

            // 또는 아래와 같이 에러 메시지를 반환할 수도 있다.
            return Err("not enough arguments");
        }
        let query = args[1].clone();
        let filename = args[2].clone();

        Ok(Config { query, filename })
    }
}

// Result<(), Box<dyn Error>> 에서 각각의 의미
// Unit type으로 ()을 반환하며 Ok 에서 이를 반환한다. 여기서 ()은 empty를 의미한다.
pub fn run(config: Config) -> Result<(), Box<dyn Error>> {
    // '?' operator를 사용하면 에러 발생 시 caller에게 처리해야할 Error 값을
    // 반환한다.
    let contents = fs::read_to_string(config.filename)?;

    println!("With text:\n{}", contents);

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn one_result() {
        let query = "duct";
        let contents = "\
Rust:
safe, fast, productive.
Pick three.";

        assert_eq!(
            vec!["safe, fast, productive."],
            search(query, contents)
        );
    }
}

pub fn search<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
    let mut results = Vec::new();

    for line in contents.lines() {
        if line.contains(query) {
            results.push(line);
        }
    }

    results
}