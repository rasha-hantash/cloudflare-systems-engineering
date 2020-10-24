# rasha

## Usage
```
    ____                                *
   /___ \_   _  ___  ___ _ __         * | *
  //  / / | | |/ _ \/ _ \ '_ \       * \|/ *
 / \_/ /| |_| |  __/  __/ | | | * * * \|O|/ * * *
 \___,_\ \__,_|\___|\___|_| |_|  \o\o\o|O|o/o/o/
  v0.1.0  by @rasha-hantash      (<><><>O<><><>)

This is a simple CLI tool that makes a request to your links page.
Requests to the page are done via sockets and not through libraries.
The program itself was designed to challenge myself with socket-level programming. 

Usage:
  go run cli.go <website you want to visit>  [flags]

Flags:
  -h, --help          help for go
  -p, --profile int   the profile flag
```

## Observations
```
I saw that the average time it took to request a workers website was much faster than requesting a popular website
```

## To Run
```
make build
./bin/main.go <website> [flags]

Example
./bin/main.go www.google.com/
./bin/main.go mini-santorini.rasha-hantash.workers.dev/links --profile 3
./bin/main.go --help
go run main.go www.youtube.com/

```
