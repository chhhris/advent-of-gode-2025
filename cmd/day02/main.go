package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("02_01_input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read all input into a single string
	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Generate all possible invalid IDs (pattern repeated twice)
	// These are numbers like 11, 22, 6464, 123123, etc.
	invalidIDs := generateInvalidIDs()

	// Parse ranges and find invalid IDs within each range
	ranges := strings.Split(strings.TrimSpace(input), ",")
	var totalSum int64 = 0

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}

		parts := strings.Split(r, "-")
		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic(err)
		}
		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic(err)
		}

		// Find all invalid IDs in this range
		for _, id := range invalidIDs {
			if id >= start && id <= end {
				totalSum += id
			}
		}
	}

	fmt.Println("Sum of all invalid IDs:", totalSum)
}

// generateInvalidIDs generates all numbers that are a pattern repeated twice
// For example: 11 (1 repeated), 6464 (64 repeated), 123123 (123 repeated)
func generateInvalidIDs() []int64 {
	var result []int64

	// Generate patterns of 1 to 5 digits
	// When repeated, this gives numbers with 2 to 10 digits
	// This covers the range needed for our input (up to ~10 billion)
	for digits := 1; digits <= 5; digits++ {
		// Calculate the range of patterns for this digit count
		// 1 digit: 1-9, 2 digits: 10-99, 3 digits: 100-999, etc.
		start := int64(1)
		for i := 1; i < digits; i++ {
			start *= 10
		}
		end := start*10 - 1

		for pattern := start; pattern <= end; pattern++ {
			// Create the repeated number by concatenating the pattern with itself
			s := strconv.FormatInt(pattern, 10)
			repeated := s + s
			id, _ := strconv.ParseInt(repeated, 10, 64)
			result = append(result, id)
		}
	}

	return result
}
