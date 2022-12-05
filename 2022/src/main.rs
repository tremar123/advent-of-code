use clap::Parser;
use std::{fs, path::PathBuf};

mod day1;
mod day2;
mod day3;
mod day4;
mod day5;

#[derive(Parser, Debug)]
struct Args {
    #[arg(short, long)]
    day: u8,

    #[arg(short, long)]
    part: u8,

    file: PathBuf,
}

fn main() {
    let args = Args::parse();

    let file = match fs::read_to_string(&args.file) {
        Ok(file) => file,
        Err(err) => {
            eprintln!("{}", err);
            std::process::exit(2);
        }
    };

    let result = match args.day {
        1 => match args.part {
            1 => day1::calorie_counter_part_1(file),
            2 => day1::calorie_counter_part_2(file),
            _ => wrong_part(),
        },
        2 => match args.part {
            1 => day2::rock_paper_scissors_part_1(file),
            2 => day2::rock_paper_scissors_part_2(file),
            _ => wrong_part(),
        },
        3 => match args.part {
            1 => day3::rucksack_reorganization_part_1(file),
            2 => day3::rucksack_reorganization_part_2(file),
            _ => wrong_part(),
        },
        4 => match args.part {
            1 => day4::camp_cleanup_part_1(file),
            2 => day4::camp_cleanup_part_2(file),
            _ => wrong_part(),
        },
        5 => match args.part {
            1 => day5::supply_stack_part_1(file),
            2 => day5::supply_stack_part_2(file),
            _ => wrong_part(),
        },
        _ => String::from("Wrong day"),
    };

    println!("{}", result);
}

fn wrong_part() -> String {
    String::from("Wrong part!")
}
