pub fn rucksack_reorganization_part_1(input: String) -> String {
    input
        .lines()
        .map(|line| {
            let halfs = line.split_at(line.len() / 2);
            for ch in halfs.0.chars() {
                if halfs.1.contains(ch) {
                    if ch.is_lowercase() {
                        return ch as u32 - 96;
                    } else {
                        return ch as u32 - 38;
                    }
                }
            }
            0
        })
        .sum::<u32>()
        .to_string()
}

pub fn rucksack_reorganization_part_2(input: String) -> String {
    let lines = input.lines().collect::<Vec<_>>();
    let mut i = 0;

    let mut results = vec![];

    while i < lines.len() {
        for ch in lines[i].chars() {
            if lines[i + 1].contains(ch) && lines[i + 2].contains(ch) {
                if ch.is_lowercase() {
                    results.push(ch as u32 - 96);
                    break;
                } else {
                    results.push(ch as u32 - 38);
                    break;
                }
            }
        }
        i += 3;
    }

    results.iter().sum::<u32>().to_string()
}
