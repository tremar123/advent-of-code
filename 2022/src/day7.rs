#![allow(dead_code)]

struct File<'a> {
    name: &'a str,
    size: u32,
}

struct Directory<'a> {
    name: &'a str,
    size: u32,
    files: Vec<File<'a>>,
    directories: Vec<Self>,
    up_dir: Option<&'a mut Self>,
}

impl<'a> Directory<'a> {
    fn change_dir(self: &mut Self, arg: &str) {
        let dir_index = self
            .directories
            .iter()
            .position(|dir| dir.name == arg)
            .unwrap();
        let _ = &mut self.directories[dir_index];
    }

    fn add_dir(self: &'a mut Self, name: &'a str) {
        let ha = Directory {
            name,
            size: 0,
            files: vec![],
            directories: vec![],
            up_dir: Some(self),
        };

        self.directories.push(ha);
    }
}

pub fn no_space_left_on_device_part_1(input: String) -> String {
    let mut root_dir = Directory {
        name: "/",
        size: 0,
        files: vec![],
        directories: vec![],
        up_dir: None,
    };

    let mut current_dir = &mut root_dir;

    let mut index: usize = 0;

    let lines = input
        .lines()
        .skip(1)
        .map(|line| line.split(" ").collect::<Vec<_>>())
        .collect::<Vec<_>>();

    while index < lines.len() {
        if lines[index][1] == "ls" {
            let mut ls_index = 1;
            while lines[index + ls_index][0] != "$" {
                if let Ok(size) = lines[index][0].parse::<u32>() {
                    current_dir.files.push(File {
                        name: lines[index][1],
                        size,
                    });
                    current_dir.size += size;
                } else {
                    // current_dir.add_dir(lines[index][2]);
                }
                ls_index += 1;
            }
        } else if lines[index][1] == "cd" {
            if lines[index][2] == ".." {
            } else {
                current_dir.change_dir(lines[index][2]);
            }
        }
        index += 1;
    }

    "its hard".to_string()
}
