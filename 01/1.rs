use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let filename = "01/input.txt";

    if let Ok(lines) = read_lines(filename) {
        let mut last = 0;
        let mut count = 0;
        for line in lines {
            if let Ok(s) = line {
                let i = s.parse::<i32>().unwrap();
                
                if  last != 0 && i>last {
                    count += 1;
                } 
                last = i
            }
        }
        println!("{}", count)
    }
}


fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
