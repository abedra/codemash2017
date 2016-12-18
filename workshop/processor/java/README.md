# Java Processor Example

This example parses the sample application log file to determine bad
actors. It blacklists the actors in the Repsheet cache.

## Setup

#### apt based systems (Ubuntu and family)

```
$ sudo apt-get install openjdk-8-jdk maven
```

#### rpm based systems (CentOS and family)

```
$ sudo yum install java-1.8.0-openjdk maven
```

#### FreeBSD

```
$ sudo pkg install openjdk8-8.102.14_1 maven33-3.3.9
```

#### macOS

This example requires Java 8. You will need to download the installer
from
[Oracle's Website](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html)

```
$ brew install maven
```

## Build and Run

```
$ mvn compile
$ mvn exec:java -Dexec.mainClass="Processor" -Dexec.args="../../app/logs/app.log 10"
Blacklisting 1.1.1.1. Failed Login Threshold 10, Actual: 31
Blacklisting 2.2.2.2. Failed Login Threshold 10, Actual: 31
Blacklisting 127.0.0.1. Failed Login Threshold 10, Actual: 31
```
