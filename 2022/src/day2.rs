const ROCK_POINTS: i32 = 1;
const PAPER_POINTS: i32 = 2;
const SCISSORS_POINTS: i32 = 3;

const LOST_POINTS: i32 = 0;
const DRAW_POINTS: i32 = 3;
const WIN_POINTS: i32 = 6;

/*
A, X - Rock
B, Y - Paper
C, Z - Scissors
*/
pub fn rock_paper_scissors_part_1(input: String) -> String {
    count_points_map(input, &|opponent, me| match (opponent, me) {
        ('A', 'X') => DRAW_POINTS + ROCK_POINTS,
        ('A', 'Y') => WIN_POINTS + PAPER_POINTS,
        ('A', 'Z') => LOST_POINTS + SCISSORS_POINTS,
        ('B', 'X') => LOST_POINTS + ROCK_POINTS,
        ('B', 'Y') => DRAW_POINTS + PAPER_POINTS,
        ('B', 'Z') => WIN_POINTS + SCISSORS_POINTS,
        ('C', 'X') => WIN_POINTS + ROCK_POINTS,
        ('C', 'Y') => LOST_POINTS + PAPER_POINTS,
        ('C', 'Z') => DRAW_POINTS + SCISSORS_POINTS,
        _ => 0,
    })
}

/*
X - LOSE
Y - DRAW
Z - WIN
*/
pub fn rock_paper_scissors_part_2(input: String) -> String {
    count_points_map(input, &|opponent, me| match (opponent, me) {
        ('A', 'X') => LOST_POINTS + SCISSORS_POINTS,
        ('A', 'Y') => DRAW_POINTS + ROCK_POINTS,
        ('A', 'Z') => WIN_POINTS + PAPER_POINTS,
        ('B', 'X') => LOST_POINTS + ROCK_POINTS,
        ('B', 'Y') => DRAW_POINTS + PAPER_POINTS,
        ('B', 'Z') => WIN_POINTS + SCISSORS_POINTS,
        ('C', 'X') => LOST_POINTS + PAPER_POINTS,
        ('C', 'Y') => DRAW_POINTS + SCISSORS_POINTS,
        ('C', 'Z') => WIN_POINTS + ROCK_POINTS,
        _ => 0,
    })
}

fn count_points_map(input: String, match_statement: &dyn Fn(char, char) -> i32) -> String {
    input
        .lines()
        .map(|round| {
            let opponent = round.chars().nth(0).unwrap();
            let me = round.chars().nth(2).unwrap();
            match_statement(opponent, me)
        })
        .sum::<i32>()
        .to_string()
}
