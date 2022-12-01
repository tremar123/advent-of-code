use std::{fs, path::PathBuf};

use clap::Parser;

mod day1;

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
            1 => day1::calorie_counter_part1(file),
            2 => day1::calorie_counter_part2(file),
            _ => String::from("Wrong puzzlec"),
        },
        _ => String::from("Wrong day"),
    };

    println!("{}", result);
}
