[![Go](https://github.com/abeltay/advent-of-code-2023/actions/workflows/go.yml/badge.svg)](https://github.com/abeltay/advent-of-code-2023/actions/workflows/go.yml)

# Advent of Code 2023

My attempt at <https://adventofcode.com/2023>

> Note: All task input files are excluded from the repository with ".gitignore" – we should not post them publicly, as Eric Wastl asks for: [Tweet](https://twitter.com/ericwastl/status/1465805354214830081).

# To run the scripts
1. Run `make DAY=01 prep` with the day number. The example here is for Day 01
2. Run `make input` to download the input file into the day's folder
3. Run `make test` to run current day's file

## To set up cookie for `make input`
1. Login to [AoC](https://adventofcode.com/)
2. Most requests should use your session cookie, copy it. It starts with "session=" until ";"
3. Put it into "cookie.txt"

## To test the template
1. Set "day.txt" to "template"
2. Run `make test`
