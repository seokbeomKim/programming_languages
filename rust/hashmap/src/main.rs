use std::collections::HashMap;

fn main() {
    let mut scores = HashMap::new();

    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);

    println!("blue: {}", scores["Blue"]);

    // 또 다른 할당 방법은 collect 메서드를 이용하여 튜플로 구성된 벡터를
    // hashmap으로 구성하는 방법이다.
    let teams = vec![String::from("Blue"), String::from("Yellow")];
    let initial_scores = vec![10, 50];

    let scores: HashMap<_, _> = teams.iter().zip(initial_scores.iter()).collect();

    let field_name = String::from("Favorite color");
    let field_value = String::from("Blue");

    let mut map = HashMap::new();
    map.insert(&field_name, &field_value);

    // 만약 field_value 를 1 로 바꿔 기존 ownership을 날려버려도, hashmap은
    // 그대로인가?
    let field_value = 1;
    println!("field_name: {}, field_value: {}", field_name, field_value);
    println!("map value: {}", map[&field_name]);

    // 그대로이다. 내부적으로 copy 가 일어났을 것으로 예상한다.
    for (key, value) in &scores {
        println!("{}: {}", key, value);
    }

    panic!();
}


