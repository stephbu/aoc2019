# Advent Of Code 2019 - Golang
I'll try to squeeze in code every day
https://adventofcode.com/2019

| Day    | Solution Code                        | Notes |
|--------|--------------------------------------|-------|
| Day 1  | <a href="./day1/">Day 1 Code</a>     |Good warmup exercise           |
| Day 2  | <a href="./day2/">Day 2 Code</a>     |Interesting problem that had a pseudo-CPU instruction processor/emulator Go functions made the solution quite interesting|
| Day 3  | <a href="./day3/">Day 3 Code</a>     |Computationally intensive tortoise/hare path enumerator - the brute force approach made Part 2 pretty easy|
| Day 4  | <a href="./day4/">Day 4 Code</a>     |Probably the trickiest to-date because state-machine needed additional test cases to catch some interesting corner cases|

## Getting Started
I've broken out the project into packages of code and tests per day.  Clone the repo, then use 

```
go get github.com/stephbu/aoc2019
cd $GOROOT/src/github.com/stephbu/aoc2019
dep ensure       
go test ./... -v 
``` 

I've taken writing unit-tests to assert samples initially, then separate tests per "part" of each challenge.