# Go Processor Example

This example parses the sample application log file to determine bad
actors. It blacklists the actors in the Repsheet cache.

## Setup

You will already have go installed and setup as part of completing the
workshop, but you will need to the following package installed

```
$ go get github.com/fzzy/radix/redis
```

## Build and Run

```
$ go build processor.go
$ ./processor -file ../../app/logs/app.log
Blacklisting 127.0.0.1. Threshold: 10, Actual: 31
Blacklisting 1.1.1.1. Threshold: 10, Actual: 31
Blacklisting 2.2.2.2. Threshold: 10, Actual: 31
```
