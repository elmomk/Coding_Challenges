use clap::Parser;
use std::fs;
use std::io;
use std::io::Read;

/// Simple program to greet a person
#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    /// file to read optional
    #[arg(short, long, default_value = "")]
    name: String,

    /// bytes flag default is False
    #[arg(short, long)]
    bytes: bool,

    /// lines flag default is False
    #[arg(short, long)]
    lines: bool,

    /// words flag default is False
    #[arg(short, long)]
    words: bool,
}

fn main() {
    let arguments = Args::parse();
    println!("{:?}", arguments);

    let content = if !arguments.name.is_empty() {
        println!("stdin is empty");
        println!("{:?}", arguments.name);
        fs::read_to_string(arguments.name).expect("Unable to read file")
    } else {
        println!("stdin is not empty");
        let stdin = io::stdin();
        let mut stdin = stdin.lock();
        let mut line = String::new();
        stdin
            .read_to_string(&mut line)
            .unwrap_or_else(|err| panic!("Unable to read file: {}", err));
        line
    };

    match (arguments.bytes, arguments.lines, arguments.words) {
        (true, false, false) => {
            let byte_count = content.as_bytes().len();
            println!("{}", byte_count);
        }
        (false, true, false) => {
            let line_count = content.lines().count();
            println!("{}", line_count);
        }
        (false, false, true) => {
            let word_count = content.split_whitespace().count();
            println!("{}", word_count);
        }
        _ => {
            let byte_count = content.as_bytes().len();
            let line_count = content.lines().count();
            let word_count = content.split_whitespace().count();
            println!("{} {} {}", byte_count, line_count, word_count);
        }
    }
}
