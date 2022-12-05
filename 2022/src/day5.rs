type Stack = Vec<char>;
type Stacks = Vec<Stack>;

struct Instructions {
    count: usize,
    from: usize,
    to: usize,
}

impl Instructions {
    fn new(line: &str) -> Self {
        let inst = line.split(" ").collect::<Vec<_>>();
        Self {
            count: inst[1].parse().unwrap(),
            from: inst[3].parse::<usize>().unwrap() - 1,
            to: inst[5].parse::<usize>().unwrap() - 1,
        }
    }
}

pub fn supply_stack_part_1(input: String) -> String {
    let (mut stacks, moves) = organize_crates_and_read_instructions(&input);

    moves.iter().for_each(|inst| {
        for _ in 0..inst.count {
            let item = stacks[inst.from].pop().unwrap();
            stacks[inst.to].push(item);
        }
    });

    last_elements_of_stacks(stacks)
}

pub fn supply_stack_part_2(input: String) -> String {
    let (mut stacks, moves) = organize_crates_and_read_instructions(&input);

    moves.iter().for_each(|inst| {
        let mut items = vec![];
        for _ in 0..inst.count {
            items.push(stacks[inst.from].pop().unwrap());
        }

        items
            .into_iter()
            .rev()
            .for_each(|item| stacks[inst.to].push(item));
    });

    last_elements_of_stacks(stacks)
}

fn organize_crates_and_read_instructions(input: &String) -> (Stacks, Vec<Instructions>) {
    let (stacks_input, moves) = input.split_once("\n\n").unwrap();

    let mut indexes = vec![];

    stacks_input
        .lines()
        .rev()
        .next()
        .unwrap()
        .chars()
        .enumerate()
        .for_each(|(index, ch)| {
            if ch.is_digit(10) {
                indexes.push(index);
            }
        });

    let mut stacks: Stacks = std::iter::repeat(vec![]).take(indexes.len()).collect();

    stacks_input.lines().rev().skip(1).for_each(|line| {
        line.chars().enumerate().for_each(|(index, ch)| {
            if ch.is_alphabetic() {
                let element_index = indexes.iter().position(|&x| x == index).unwrap();
                stacks[element_index].push(ch);
            }
        })
    });

    let moves = moves.lines().map(|line| Instructions::new(line)).collect();

    (stacks, moves)
}

fn last_elements_of_stacks(stacks: Stacks) -> String {
    stacks
        .iter()
        .map(|stack| stack.last().unwrap())
        .collect::<String>()
}
