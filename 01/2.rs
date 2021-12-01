use std::{
    fs::File,
    io::{prelude::*, BufReader},
    path::Path,
};

fn main() {
    let filename = "01/input.txt";

    let lines = lines_from_file(filename);

    let v:Vec<i32> = lines.into_iter().map(|s| s.parse::<i32>().unwrap()).collect();

    let mut values = &v[..];

    let mut count = 0;
    let mut last = 0;

    while values.len() >= 3 {
        let i = sum3(values);
        if last != 0 && i > last {
            count += 1;
        }
        last = i;
        values = &values[1..]
    }
    println!("{}",count)
}

fn sum3(in_vec :&[i32]) -> i32 {
    in_vec[0] + in_vec[1]+ in_vec[2]
}


fn lines_from_file(filename: impl AsRef<Path>) -> Vec<String> {
    let file = File::open(filename).expect("no such file");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("Could not parse line"))
        .collect()
}
