# C++ Processor Example

This example parses the sample application log file to determine bad
actors. It blacklists the actors in the Repsheet cache.

## Setup

#### apt based systems (Ubuntu and family)

```
$ sudo apt-get install cmake libboost-dev libboost-system-dev libboost-filesystem-dev
  libboost-test-dev libboost-date-time-dev libboost-regex-dev libboost-program-options-dev
  libmsgpack-dev
```

#### rpm based systems (CentOS and family)

```
$ sudo yum install cmake boost-devel msgpack-devel
```

#### FreeBSD

```
$ sudo pkg install cmake boost-all msgpack
```

#### macOS

```
$ brew install cmake boost msgpack
```

## Build and Run

First, you will need to clone and install `redis3m`

```
$ git clone https://github.com/luca3m/redis3m
$ cd redis3m
$ cmake .
$ make
$ sudo make install
```

Next, compile and run the example. If you are on a mac you should use
the Makefile to compile. If you are on FreeBSD substitute `g++` for
`clang++`.

```
$ g++ -std=c++11 processor.cpp -o processor -lboost_program_options -lredis3m
$ ./processor -l ../../app/logs/app.log -t 10
Blacklisting 2.2.2.2. Threshold: 10, Actual: 31
Blacklisting 1.1.1.1. Threshold: 10, Actual: 31
Blacklisting 127.0.0.1. Threshold: 10, Actual: 31
```
