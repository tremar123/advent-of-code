pub fn camp_cleanup_part_1(input: String) -> String {
    input
        .lines()
        .filter(|line| {
            let ranges = line.split(",").collect::<Vec<_>>();
            let first = ranges[0].split("-").collect::<Vec<_>>();
            let second = ranges[1].split("-").collect::<Vec<_>>();

            let first = (
                first[0].parse::<i32>().unwrap(),
                first[1].parse::<i32>().unwrap(),
            );

            let second = (
                second[0].parse::<i32>().unwrap(),
                second[1].parse::<i32>().unwrap(),
            );

            if first.0 >= second.0 && first.1 <= second.1
                || second.0 >= first.0 && second.1 <= first.1
            {
                return true;
            }
            return false;
        })
        .count()
        .to_string()
}

pub fn camp_cleanup_part_2(input: String) -> String {
    input
        .lines()
        .filter(|line| {
            let ranges = line.split(",").collect::<Vec<_>>();
            let first = ranges[0].split("-").collect::<Vec<_>>();
            let second = ranges[1].split("-").collect::<Vec<_>>();

            let first = (
                first[0].parse::<i32>().unwrap(),
                first[1].parse::<i32>().unwrap(),
            );

            let second = (
                second[0].parse::<i32>().unwrap(),
                second[1].parse::<i32>().unwrap(),
            );

            if first.1 >= second.0 && first.1 <= second.1
                || first.0 <= second.1 && first.0 >= second.0
                || second.0 >= first.0 && second.0 <= first.1
                || second.1 >= first.0 && second.1 <= first.1
            {
                return true;
            }
            return false;
        })
        .count()
        .to_string()
}
