use std::thread;
use std::time::Duration;

fn simulated_expensive_calculation(intensity: u32) -> u32 {
    println!("calculating slowly...");
    thread::sleep(Duration::from_secs(2));
    intensity
}

fn generate_workout(intensity: u32, random_number: u32) {
    if intensity < 25 {
        println!(
            "Today, do {} pushups!",
            simulated_expensive_calculation(intensity)
        );
        println!(
            "Next, do {} situps!",
            simulated_expensive_calculation(intensity)
        );
    }
}

fn main() {
    let simulated_user_specified_value = 10;
    let simulated_random_number = 7;
    let example_closure = |x: u32| x;
    let y = example_closure(10);

    generate_workout(
        simulated_user_specified_value,
        simulated_random_number
    );
    simulated_expensive_calculation(2);
}
