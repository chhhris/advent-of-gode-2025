package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day_1_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	position := 50
	zeroCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if direction == 'R' {
			position = (position + distance) % 100
		} else { // 'L'
			position = (position - distance) % 100
			if position < 0 {
				position += 100
			}
		}

		if position == 0 {
			zeroCount++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Password (times dial points at 0):", zeroCount)
}

