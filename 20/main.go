package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abeltay/advent-of-code-2023/utilities"
)

func parseFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var arr []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		arr = append(arr, t)
	}
	return arr
}

type moduleType int

const (
	broadcast moduleType = iota
	flipFlop
	conjunction
)

type pulse struct {
	from        string
	isPulseHigh bool
	to          string
}

type module struct {
	name      string
	kind      moduleType
	ffState   bool
	conjunc   []pulse
	broadcast []string
}

func (f *module) receive(in pulse) []pulse {
	switch f.kind {
	case flipFlop:
		if in.isPulseHigh {
			return []pulse{}
		} else {
			f.ffState = !f.ffState
			in.isPulseHigh = f.ffState
		}
	case conjunction:
		for i := range f.conjunc {
			if f.conjunc[i].from == in.from {
				f.conjunc[i].isPulseHigh = in.isPulseHigh
				break
			}
		}
		var foundLow bool
		for i := range f.conjunc {
			if !f.conjunc[i].isPulseHigh {
				foundLow = true
				break
			}
		}
		in.isPulseHigh = foundLow
	default:
	}
	out := make([]pulse, len(f.broadcast))
	for i := range f.broadcast {
		out[i] = pulse{
			from:        f.name,
			to:          f.broadcast[i],
			isPulseHigh: in.isPulseHigh,
		}
	}
	return out
}

func parseInput(filename string) []module {
	in := parseFile(filename)

	modules := []module{}
	for _, row := range in {
		s := strings.Split(row, " -> ")
		var kind moduleType
		name := s[0]
		switch name[0] {
		case '%':
			kind = flipFlop
		case '&':
			kind = conjunction
		default:
			kind = broadcast
		}
		if kind != broadcast {
			name = name[1:]
		}
		modules = append(modules, module{
			name:      name,
			kind:      kind,
			broadcast: strings.Split(s[1], ", "),
		})
	}
	for current := range modules {
		if modules[current].kind == conjunction {
			for other := range modules {
				for _, dest := range modules[other].broadcast {
					if dest == modules[current].name {
						modules[current].conjunc = append(modules[current].conjunc, pulse{from: modules[other].name})
					}
				}
			}
		}
	}
	return modules
}

func part1(filename string) int {
	in := parseInput(filename)
	var totalLow, totalHigh int
	for i := 0; i < 1000; i++ {
		pulses := []pulse{
			{from: "button", to: "broadcaster"},
		}
		for len(pulses) > 0 {
			for _, v := range pulses {
				if v.isPulseHigh {
					totalHigh++
				} else {
					totalLow++
				}
			}
			var newPulses []pulse
			for _, v := range pulses {
				for cur := range in {
					if in[cur].name == v.to {
						newP := in[cur].receive(v)
						newPulses = append(newPulses, newP...)
					}
				}
			}
			pulses = newPulses
		}
	}
	return totalLow * totalHigh
}

func lcm(multiples [][]int) int {
	factor := make([]int, len(multiples))
	for i := range multiples {
		factor[i] = multiples[i][1] - multiples[i][0]
	}
	return utilities.LCM(factor[0], factor[1], factor[2:]...)
}

func part2(filename string) int {
	in := parseInput(filename)
	i := 0
	multiples := make([][]int, 4)
	for {
		pulses := []pulse{
			{from: "button", to: "broadcaster"},
		}
		for len(pulses) > 0 {
			var newPulses []pulse
			for _, v := range pulses {
				for cur := range in {
					if in[cur].name == v.to {
						newP := in[cur].receive(v)
						newPulses = append(newPulses, newP...)
					}
				}
			}
			for _, v := range newPulses {
				if v.isPulseHigh {
					switch v.from {
					case "rz":
						multiples[0] = append(multiples[0], i)
					case "lf":
						multiples[1] = append(multiples[1], i)
					case "br":
						multiples[2] = append(multiples[2], i)
					case "fk":
						multiples[3] = append(multiples[3], i)
					default:
						continue
					}
					var notDone bool
					for _, v := range multiples {
						if len(v) < 3 {
							notDone = true
						}
					}
					if !notDone {
						return lcm(multiples)
					}
				}
			}
			pulses = newPulses
		}
		i++
	}
}
