package main

import (
	"advent-of-code-2022/file-reader"
	"fmt"
	"strconv"
)

func main() {
	var content, _ = file_reader.ReadFile("day1.txt")

	var highest int64 = 0
	var total int64 = 0

	for _, line := range content {
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

	fmt.Println(highest)
}
