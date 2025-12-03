use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

const NUMBER_OF_DIGITS: usize = 12;

fn main() -> io::Result<()> {
    let path = Path::new("input.txt");
    let file = File::open(path)?;
    let reader = io::BufReader::new(file);
    let mut sum: u64 = 0;

    for line in reader.lines() {
        let line = line?;
        let digits = convert_string_to_digits(&line);
        let highest_digits = find_highest_digits(&digits, NUMBER_OF_DIGITS);

        for (i, digit) in highest_digits.iter().enumerate() {
            sum += *digit as u64 * 10_u64.pow((NUMBER_OF_DIGITS - i - 1) as u32)
        }
    }

    println!("{}", sum);

    Ok(())
}

fn convert_string_to_digits(s: &String) -> Vec<u8> {
    s.bytes().map(|b| b - b'0').collect()
}

fn find_highest_digits(digits: &Vec<u8>, number_of_digits: usize) -> Vec<u8> {
    let mut start = 0;
    let mut end = digits.len() - (number_of_digits - 1);
    let mut result = Vec::new();

    for _i in 0..number_of_digits {
        let number = find_max_of_vector(&digits[start..end].to_vec());

        start += &digits[start..end]
            .iter()
            .position(|&x| x == number)
            .unwrap()
            + 1;
        end += 1;

        result.push(number);
    }

    result
}

fn find_max_of_vector(digits: &Vec<u8>) -> u8 {
    *digits.iter().max().unwrap()
}
