package main

import (
	"advent-of-code-2022/file-reader"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var content, _ = file_reader.ReadFile("day1.txt")

	var highest int64 = 0
	var total int64 = 0

	for index, line := range content {
		if line != "" && index == len(content)-1 {
			var value, _ = strconv.ParseInt(line, 10, 32)

			total += value

			if total > highest {
				highest = total
			}

			continue
		}

		if line == "" {
			if total > highest {
				highest = total
			}

			total = 0

			continue
		}

		var value, _ = strconv.ParseInt(line, 10, 32)
		total += value
	}

	topThreeSum := getTop3Sum(content)

	fmt.Println(highest)
	fmt.Print(topThreeSum)
}

type Elf struct {
	total int64
}

func getTop3Sum(content []string) int64 {
	var elves []Elf

	var total int64 = 0

	for index, line := range content {
		if line != "" && index == len(content)-1 {
			var value, _ = strconv.ParseInt(line, 10, 32)

			total += value

			elves = append(elves, Elf{total: total})

			continue
		}

		if line == "" {
			elves = append(elves, Elf{total: total})
			total = 0

			continue
		}

		var value, _ = strconv.ParseInt(line, 10, 32)
		total += value
	}

	sort.Slice(elves, func(p, q int) bool {
		return elves[p].total > elves[q].total
	})

	return elves[0].total + elves[1].total + elves[2].total
}
