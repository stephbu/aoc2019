# Advent Of Code 2019 - Golang
I'll try to squeeze in code every day
https://adventofcode.com/2019

|Day|Source Package|Puzzle Rubric|Notes|
|---|---|---|---|
|1|<a href="./day1/">day1</a>|<a href="https://adventofcode.com/2019/day/1">day1</a>|Good warmup exercise|
|2|<a href="./day2/">day2</a>|<a href="https://adventofcode.com/2019/day/2">day2</a>|Interesting problem that had a pseudo-CPU instruction processor/emulator Go functions made the solution quite interesting|
|3|<a href="./day3/">day3</a>|<a href="https://adventofcode.com/2019/day/3">day3</a>|Computationally intensive tortoise/hare path enumerator - the brute force approach made Part 2 pretty easy|
|4|<a href="./day4/">day4</a>|<a href="https://adventofcode.com/2019/day/4">day4</a>|Probably the trickiest to-date because state-machine needed additional test cases to catch some interesting corner cases|
|5|<a href="./day5/">day5</a>|<a href="https://adventofcode.com/2019/day/5">day5</a>|The investment in writing a pseudo-CPU on day 2 just paid back, adding instructions and operating modes was pretty easy, albeit needed a few test cases for things like the parameter masking|

## Getting Started
I've broken out the project into packages of code and tests per day.  Clone the repo, then use 

```
go get github.com/stephbu/aoc2019
cd $GOROOT/src/github.com/stephbu/aoc2019
dep ensure       
go test ./... -v   # test all days
...
go test ./day4 -v  # test specific day
``` 

I've taken writing unit-tests to assert samples initially, then separate tests per "part" of each challenge.