pub fn calorie_counter_part_1(input: String) -> String {
    count_carried(input).iter().max().unwrap().to_string()
}

pub fn calorie_counter_part_2(input: String) -> String {
    count_carried(input)
        .into_iter()
        .fold([0; 3], |mut acc, num| {
            if num > acc[0] {
                acc[2] = acc[1];
                acc[1] = acc[0];
                acc[0] = num;
            } else if num > acc[1] {
                acc[2] = acc[1];
                acc[1] = num;
            } else if num > acc[2] {
                acc[2] = num;
            }
            acc
        })
        .iter()
        .sum::<i32>()
        .to_string()
}

fn count_carried(input: String) -> Vec<i32> {
    input
        .split("\n\n")
        .map(|elf| {
            elf.split("\n")
                .map(|num| match num.parse::<i32>() {
                    Ok(x) => x,
                    Err(_) => 0,
                })
                .sum()
        })
        .collect::<Vec<i32>>()
}
