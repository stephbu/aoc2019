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
|6|<a href="./day6/">day6</a>|<a href="https://adventofcode.com/2019/day/6">day6</a>|Stalled for a while because I missed the test data/samples being ordered, but unordered in the puzzle. This caused problems in missing large chunks of data during the expansion of direct orbits into indirect orbits, and undercounting the dataset (~7300 vs ~203000).  Rectified the mistake by taking two passes at the data, one to load the direct orbits, and second to recursively scan and expand each indirect orbit.  Part two just needed to retro-fit hop-counting in the indirect expansion phase, to enable fast finding of common-root brute-force list comparison to extract least hops|
|7|<a href="./day7/">day7</a>|<a href="https://adventofcode.com/2019/day/7">day7</a>|Third day of Intcode CPU, after starting Part 1 a little late, I made good progress and took less than 30mins to get to answer.  Part 2 not so good, managed to miss the fact that I was using a ref-copy of memory, so go-routines were corrupting each other.  Figured it out in the end.|
## Getting Started
I've broken out the project into packages of code and tests per day.  Clone the repo, then use 

```
go get github.com/stephbu/aoc2019
cd $GOROOT/src/github.com/stephbu/aoc2019
dep ensure       
go test ./... -v   # test all days
...
go test ./day6 -v  # test specific day
``` 

I've taken writing unit-tests to assert samples initially, then separate tests per "part" of each challenge.