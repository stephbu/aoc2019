# Advent Of Code 2019 - Golang
I'll try to squeeze in code every day
https://adventofcode.com/2019

| Day    | Solution Code                        |
|--------|--------------------------------------|
| Day 1  | <a href="./day1/">Day 1 Code</a>     |
| Day 2  | <a href="./day2/">Day 2 Code</a>     |

## Getting Started
I've broken out the project into packages of code and tests per day.  Clone the repo, then use 

```
**[~]$**        go get github.com/stephbu/aoc2019
**[~]$**        cd $GOROOT/src/github.com/stephbu/aoc2019
**[aoc2019]$**: dep ensure       ## pull down vendored packages
**[aoc2019]$**: go test ./... -v ## Run all tests
``` 

I've taken writing unit-tests to assert samples initially, then separate tests per "part" of each challenge.