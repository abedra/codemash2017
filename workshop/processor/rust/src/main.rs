use std::collections::HashMap;
use std::path::Path;
use std::fs::File;
use std::error::Error;
use std::io::BufReader;
use std::io::prelude::*;
use std::collections::hash_map::Entry::{Vacant, Occupied};

fn process(entries: &mut HashMap<String, i64>, line: &str) {
    let parts: Vec<&str> = line.split(' ').collect();
    let address = parts[0].to_string();
    let method = parts[5].replace("\"", "").to_string();
    let response_code = parts[8].to_string();

    if method == "POST" && response_code == "200" {
        match entries.entry(address) {
            Vacant(entry) => { entry.insert(1); },
            Occupied(mut entry) => { *entry.get_mut() += 1; },
        }
    }
}

fn main() {
    let path = Path::new("../../app/logs/app.log");
    let file_handle = match File::open(&path) {
        Err(why) => {
            println!("Could not open {} : {}", path.display(), Error::description(&why));
            return;
        },
        Ok(file) => file,
    };

    let reader = BufReader::new(file_handle);
    let mut entries = HashMap::<String, i64>::new();
    for line in reader.lines() {
        match line {
            Ok(line) => process(&mut entries, &line),
            Err(e)   => println!("ERROR {}", e),
        }
    }

    if entries.len() > 0 {
        for (address, value) in entries.iter() {
            println!("Blacklisting {}. Actual {}, Threshold {}", address, value, 10);
        }
    }
}
