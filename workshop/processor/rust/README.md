# Rust Processor Example

This example parses the sample application log file to determine bad
actors. It blacklists the actors in the Repsheet cache.

## Setup

You will need to have rust and cargo installed. The easiest way to
complete this is to use the script directly from
the [rust website](https://www.rust-lang.org/en-US/downloads.html).

```
$ curl -sSf https://static.rust-lang.org/rustup.sh | sh
```

## Build and Run

```
$ cargo run processor
    Finished debug [unoptimized + debuginfo] target(s) in 0.0 secs
     Running `target/debug/processor processor`
Blacklisting 2.2.2.2. Actual 31, Threshold 10
Blacklisting 127.0.0.1. Actual 31, Threshold 10
Blacklisting 1.1.1.1. Actual 31, Threshold 10
```
