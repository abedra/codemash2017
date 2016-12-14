# .NET Core Processor Example

This example parses the sample application log file to determine bad
actors. It blacklists the actors in the Repsheet cache.

## Setup

You will need to have .NET core installed on your machine. To install
please follow
the [setup instructions](https://www.microsoft.com/net/core) for
your machine.

## Build and Run

```
$ dotnet restore
$ dotnet build
$  dotnet run -- --file ../../app/logs/app.log --threshold 10
Blacklisting 127.0.0.1. Threshold 10, Actual 31
Blacklisting 2.2.2.2. Threshold 10, Actual 31
Blacklisting 1.1.1.1. Threshold 10, Actual 31
```
