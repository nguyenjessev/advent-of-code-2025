use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() -> io::Result<()> {
    let path = Path::new("input.txt");
    let file = File::open(path)?;
    let reader = io::BufReader::new(file);
    let mut sum: i32 = 0;

    for line in reader.lines() {
        let line = line?;
        let digits = convert_string_to_digits(line);
        let (first_number, second_number) = find_highest_pair(&digits);

        sum += first_number as i32 * 10;
        sum += second_number as i32;
    }

    println!("{}", sum);

    Ok(())
}

fn convert_string_to_digits(s: String) -> Vec<u8> {
    s.bytes().map(|b| b - b'0').collect()
}

fn find_highest_pair(digits: &Vec<u8>) -> (u8, u8) {
    let first_number = find_max_of_vector(&digits[0..digits.len() - 1].to_vec());
    let index_of_first_number = digits.iter().position(|&x| x == first_number).unwrap();
    let second_number = find_max_of_vector(&digits[index_of_first_number + 1..].to_vec());

    (first_number, second_number)
}

fn find_max_of_vector(digits: &Vec<u8>) -> u8 {
    *digits.iter().max().unwrap()
}
