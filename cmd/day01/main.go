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

		// Count how many times we pass through or land on 0 during this rotation
		var zerosInRotation int
		if direction == 'R' {
			// For right rotations: count how many times we wrap past 99->0
			zerosInRotation = (position + distance) / 100
			position = (position + distance) % 100
		} else { // 'L'
			// For left rotations: count how many times we wrap past 0->99
			if position == 0 {
				zerosInRotation = distance / 100
			} else if distance >= position {
				zerosInRotation = 1 + (distance-position)/100
			} else { // distance < position
				zerosInRotation = 0
			}
			position = (position - distance) % 100
			if position < 0 {
				position += 100
			}
		}

		zeroCount += zerosInRotation
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Password (method 0x434C49434B):", zeroCount)
}

