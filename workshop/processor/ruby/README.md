# Ruby Processor Example

This example parses the sample application log file to determine bad
actors. It blacklists the actors in the Repsheet cache.

## Setup

You will need to install [RVM](https://rvm.io/rvm/install)
or [rbenv](https://github.com/rbenv/rbenv). You will also need bundler
installed if it's not already available.

```
$ gem install bundler
$ bundle
```

## Build and Run

```
$ ruby processor.rb ../../app/logs/app.log 10
Blacklisting 127.0.0.1. Threshold 10, Actual: 31
Blacklisting 1.1.1.1. Threshold 10, Actual: 31
Blacklisting 2.2.2.2. Threshold 10, Actual: 31
```
