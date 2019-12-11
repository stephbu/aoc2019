# Advent Of Code 2019 - Golang
I'll try to squeeze in code every day
https://adventofcode.com/2019

Because each "day" is published at 9pm PST, there's a crush on quality vs. sleep and a day-job. I'm primarily focused deepening my Go skills while "getting 'er done" with crude beauty :).  I've taken writing unit-tests to assert samples initially, then separate tests per "part" of each challenge.

LMK if you're enjoying or loathing pawing over my code and approach: <a href="mailto:stephbu@gmail.com">stephbu@gmail.com</a>

|Day|Source Package|Puzzle Rubric|Notes|
|---|---|---|---|
|1|<a href="./day1/">day1</a>|<a href="https://adventofcode.com/2019/day/1">day1</a>|Good warmup exercise|
|2|<a href="./day2/">day2</a>|<a href="https://adventofcode.com/2019/day/2">day2</a>|Interesting problem that had a pseudo-CPU instruction processor/emulator Go functions made the solution quite interesting|
|3|<a href="./day3/">day3</a>|<a href="https://adventofcode.com/2019/day/3">day3</a>|Computationally intensive tortoise/hare path enumerator - the brute force approach made Part 2 pretty easy|
|4|<a href="./day4/">day4</a>|<a href="https://adventofcode.com/2019/day/4">day4</a>|Probably the trickiest to-date because state-machine needed additional test cases to catch some interesting corner cases|
|5|<a href="./day5/">day5</a>|<a href="https://adventofcode.com/2019/day/5">day5</a>|The investment in writing a pseudo-CPU on day 2 just paid back, adding instructions and operating modes was pretty easy, albeit needed a few test cases for things like the parameter masking|
|6|<a href="./day6/">day6</a>|<a href="https://adventofcode.com/2019/day/6">day6</a>|Stalled for a while because I missed the test data/samples being ordered, but unordered in the puzzle. This caused problems in missing large chunks of data during the expansion of direct orbits into indirect orbits, and undercounting the dataset (~7300 vs ~203000).  Rectified the mistake by taking two passes at the data, one to load the direct orbits, and second to recursively scan and expand each indirect orbit.  Part two just needed to retro-fit hop-counting in the indirect expansion phase, to enable fast finding of common-root brute-force list comparison to extract least hops|
|7|<a href="./day7/">day7</a>|<a href="https://adventofcode.com/2019/day/7">day7</a>|Third day of Intcode CPU, after starting Part 1 a little late, I made good progress and took less than 30mins to get to answer.  Part 2 not so good, managed to miss the fact that I was using a ref-copy of memory, so go-routines were corrupting each other.  Figured it out in the end.|
|8|<a href="./day8/">day8</a>|<a href="https://adventofcode.com/2019/day/8">day8</a>|Simpler day, easier problem, less satisfying outcome with weaker tests because of the shape fo the results require visual interpretation. I used GSheets conditional formatting to make Part 2 answer really clear.|
|9|<a href="./day9/">day9</a>|<a href="https://adventofcode.com/2019/day/9">day9</a>|More use of the IntCode Processor, adding modes for parameters working ok|
|10|<a href="./day10/">day10</a>|<a href="https://adventofcode.com/2019/day/10">day10</a>|Lots of geometry - rusty as heck.  Missed the occlusion at any angle clause, ended up writing a masking system that I canned when I RTFM.  The Hashmap/SortedList worked really well in the second part.  Most of the code was mainly assertives test cases.  Almost went to bed to sleep on Part 2, transforming the X-oriented geometry to bearings. Had a brainwave literally just before shutdown, and got 'er done.|
|11|<a href="./day11/">day11</a>|<a href="https://adventofcode.com/2019/day/11">day11</a>|Multidimensional Arrays suck lots in Golang - untuitive syntax, initialization, consumption, the works. Got there in the end.|
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

## DevLog Day 11
So I made it to the brink of day 11. On track at 20 stars and just waiting for the clock to to get another couple.  Made a tonne of use of the IntCode interpreter.  I think all those days of Z80 are starting to pay off 42 years later.  I'm sitting at 5.3KLoC, not bad for 10 nights of hacking away.
Not sure if I'll get time code Thur/Fri, I'm day-tripping SFO in a couple of days time, then company Xmas party the day after.
Pretty tired after late nights     
