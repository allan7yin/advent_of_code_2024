use std::{
    fs::File,
    io::{self, BufRead},
};

use crate::days::ChristmasSaver;

fn read_input() -> Vec<Vec<i32>> {
    let path = "src/data_files/day2.txt";

    let file = File::open(path).unwrap();
    let reader = io::BufReader::new(file);

    let mut rows: Vec<Vec<i32>> = vec![];
    for line in reader.lines() {
        let line = line.unwrap();
        let row: Vec<i32> = line
            .split_whitespace()
            .filter_map(|val| val.parse::<i32>().ok())
            .collect();

        rows.push(row);
    }

    rows
}

fn is_valid(row: &[i32]) -> bool {
    if row.len() == 1 {
        return true;
    }

    let is_increasing = row[1] > row[0];
    for i in 1..row.len() {
        let diff = (row[i] - row[i - 1]).abs();
        if diff < 1 || diff > 3 || (is_increasing != (row[i] > row[i - 1])) {
            return false;
        }
    }

    true
}

impl ChristmasSaver {
    pub fn find_safe_reports(&self) -> i32 {
        let mut safe_count = 0;
        let rows = read_input();

        for row in rows.iter() {
            if is_valid(row) {
                safe_count += 1;
            }
        }

        safe_count
    }

    pub fn find_safe_reports_tolerance(&self) -> i32 {
        let mut safe_count = 0;
        let rows = read_input();

        for row in rows.iter() {
            if is_valid(row) {
                safe_count += 1;
                continue;
            }

            let mut found_safe = false;
            for i in 0..row.len() {
                let mut temp_row = row.to_vec();
                temp_row.remove(i);
                if is_valid(&temp_row) {
                    found_safe = true;
                    break;
                }
            }

            if found_safe {
                safe_count += 1;
            }
        }

        safe_count
    }
}
