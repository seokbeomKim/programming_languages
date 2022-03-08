fn largest_i32(list: &[i32]) -> i32 {
    let mut largest = list[0];

    // for &item in list.iter() {
    //     if item > largest {
    //         largest = item;
    //     }
    // }

    // iter을 굳이 불이는 이유가 뭔가? 차이점이 있을까?
    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }

    largest
}

fn largest_char(list: &[char]) -> char {
    let mut largest = list[0];

    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }

    largest
}

fn largest<T>(list: &[T]) -> T {
    let mut largest = list[0];

    for &item in list.iter() {
        // 비교 기능을 위해 관련된 trait 을 구현해야 한다.
        if item > largest {
            largest = item;
        }
    }

    largest
}

fn main() {
    let number_list = vec![34, 50, 25, 100, 65];
    let result = largest(&number_list);
    println!("The largest number is {}", result);

    // 벡터 변경을 위해 largest number를 찾는 부분은 함수로 변경할 수 있지 않나?
    let number_list = vec![102, 34, 6000, 89, 54, 2, 43, 8];
    let result = largest_i32(&number_list);
    println!("The largest number is {}", result);
}
