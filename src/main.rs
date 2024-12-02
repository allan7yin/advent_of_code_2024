mod days;
fn main() {
    let day1 = days::ChristmasSaver {};
    println!("Day 1:");
    println!("{}", day1.get_list_distance());
    println!("{}", day1.get_similarity_score());

    let day2 = days::ChristmasSaver {};
    println!("Day 2:");
    println!("{}", day2.find_safe_reports());
    println!("{}", day2.find_safe_reports_tolerance());
}
