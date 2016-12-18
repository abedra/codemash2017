# Building a Dynamic Security Response Environment

This repository contains the slides and requisite material for the
*Building a Dynamic Security Response Environment* workshop. The
workshop will focus on setting up and building out a fully functional
web application defense system.

In order to make the most of our time together you should come
prepared. This means installing the appropriate things on your system
ahead of time. The following list of dependencies are required to
complete the exercises. If you have trouble installing any of them
there will be some time at the beginning of the workshop to sort
things out, but not long.

## An important note

There is no support for Windows in this workshop. The examples will
only work on Linux or Unix systems. If you have a windows laptop,
please install Linux inside a virtual machine or come prepared to work
with another person during the workshop.

## Pre-Workshop Setup

You will need to install the following dependencies.

* redis
* gcc
* autotools
* make
* pkgconfig
* zlib
* pcre
* check
* ruby and rubygems
* go

With the multitude of package managers out there it may not be obvious
how to best install these dependencies. This is an attempt to cover
the most common cases. If you don't see your specific setup and have
questions, please file an issue and I will try to accommodate.

There are many options for installing ruby, but the most common are
[rvm](https://rvm.io/) and
[rbenv](https://github.com/sstephenson/rbenv). You can use Ruby 1.9 or
higher for this workshop.


#### OS X

You will need to have [homebrew](http://brew.sh/) installed. Please
see their website for install instructions. Make sure you have run
`brew doctor` and resolved any issues and that you have run `brew
update` to get the latest version of any dependencies.

Installing the requisite xcode develper tools will get you most of the
dependencies listed above. Next run the following:

```sh
$ brew install redis check go automake autoconf libtool pkg-config pcre gpg
```

#### Apt based systems (Debian, Ubuntu, etc)

This particular example was done on Ubuntu 14.04 server. Your
experience may vary but it should be close.

```sh
$ sudo apt-get install build-essential autoconf automake libtool
  pkgconfig redis-server check libpcre3-dev zlib1g-dev
  libcurl4-openssl-dev golang
```

#### Rpm based systems (Redhat, Fedora, Centos, etc)

This particular example was done on Centos 7 minimal. Your experience
may vary but it should be close.

```sh
$ sudo yum groupinstall "Developmentc Tools"
$ sudo yum install epel-release
$ sudo yum install check check-devel redis pcre-devel zlib-devel curl-devel go
$ sudo systemctl start redis
```

#### FreeBSD

FreeBSD 11 was used for this reference. If you are using FreeBSD keep
in mind that all examples calling for `make` expect GNU make. You
should substitute `make` for `gmake` at all times during the workshop.

```
$ sudo pkg install git go gmake redis pkgconf automake libtool libcheck gnupg
```

#### Installing Ruby (only for running repsheet-nginx test suite)

There are a few standard ways to install Ruby. If you don't have a
preferred method and want to run the test suite you can
try [RVM](https://rvm.io/rvm/install)
or [rbenv](https://github.com/rbenv/rbenv).

#### GO environment

In order to run the examples in the workshop you will need to have a
working go setup. If you do not already have a go environment setup
you can use the `workshop_sources` file.

```
$ source workshop_sources
```

This will setup your `$GOPATH` inside of this workshop directory and
modify your `$PATH` to support any binaries installed by `go get`.

#### Processor Examples

During the course of the workshop you will be writing some code to
parse application log files. Since everyone may not get complete this
part of the workshop I have provided some examples in the `processors`
folder. Each of the examples is in a folder named after the language
it belongs to. There are individual README files in each folder that
explain how to build and run the example.
