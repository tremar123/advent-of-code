pub fn tuning_trouble_part_1(input: String) -> String {
    solve(input, 4)
}

pub fn tuning_trouble_part_2(input: String) -> String {
    solve(input, 14)
}

fn solve(input: String, length: usize) -> String {
    let buffer = input.chars().collect::<Vec<_>>();

    let mut i = 0;

    while i < buffer.len() {
        let mut j = 0;
        let mut stream = vec![];

        while j < length {
            if !stream.contains(&buffer[i + j]) {
                stream.push(buffer[i + j]);
            } else {
                break;
            }

            j += 1;
        }

        if stream.len() == length {
            return (i + j).to_string();
        }

        i += 1
    }

    "Should never get here".to_string()
}
