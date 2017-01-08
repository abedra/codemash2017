extern crate getopts;

use std::collections::HashMap;
use std::path::Path;
use std::fs::File;
use std::error::Error;
use std::io::BufReader;
use std::io::prelude::*;
use std::collections::hash_map::Entry::{Vacant, Occupied};

use getopts::Options;
use std::env;

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

fn print_usage(program: &str, opts: Options) {
    let brief = format!("Usage: {} FILE [options]", program);
    print!("{}", opts.usage(&brief));
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let program = args[0].clone();
    let mut opts = Options::new();

    opts.optopt("l", "logfile", "Path to log file", "PATH");
    opts.optopt("t", "threshold", "Blacklist threshold", "THRESHOLD");
    opts.optflag("h", "help", "Print help");

    let matches = match opts.parse(&args[1..]) {
        Ok(m) => { m }
        Err(f) => { panic!(f.to_string()) }
    };

    if matches.opt_present("h") {
        print_usage(&program, opts);
        return;
    }

    let log_file = match matches.opt_str("l") {
        Some(x) => x,
        None => "../../app/logs/app.log".to_string(),
    };

    let threshold = match matches.opt_str("t") {
        Some(x) => x.parse().unwrap(),
        None => 10
    };

    let path = Path::new(&log_file);
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
            println!("Blacklisting {}. Actual {}, Threshold {}", address, value, threshold);
        }
    }
}
