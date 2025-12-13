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

	// Generate all possible invalid IDs (pattern repeated at least twice)
	// These are numbers like 11, 111, 1111, 6464, 646464, 123123, 123123123, etc.
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
		for id := range invalidIDs {
			if id >= start && id <= end {
				totalSum += id
			}
		}
	}

	fmt.Println("Sum of all invalid IDs:", totalSum)
}

// generateInvalidIDs generates all numbers that are a pattern repeated at least twice
// For example: 11, 111, 1111 (1 repeated 2, 3, 4 times), 6464, 646464 (64 repeated 2, 3 times), etc.
func generateInvalidIDs() map[int64]bool {
	result := make(map[int64]bool)

	// Generate patterns of 1 to 5 digits
	// Maximum total digits we need is 10 (for numbers up to ~10 billion)
	for patternLen := 1; patternLen <= 5; patternLen++ {
		// Calculate the range of patterns for this digit count
		// 1 digit: 1-9, 2 digits: 10-99, 3 digits: 100-999, etc.
		patternStart := int64(1)
		for i := 1; i < patternLen; i++ {
			patternStart *= 10
		}
		patternEnd := patternStart*10 - 1

		// Maximum repetitions: 10 / patternLen (since max total digits is 10)
		maxRepetitions := 10 / patternLen

		for pattern := patternStart; pattern <= patternEnd; pattern++ {
			patternStr := strconv.FormatInt(pattern, 10)

			// Repeat the pattern 2 or more times
			for reps := 2; reps <= maxRepetitions; reps++ {
				var repeated strings.Builder
				for r := 0; r < reps; r++ {
					repeated.WriteString(patternStr)
				}
				id, _ := strconv.ParseInt(repeated.String(), 10, 64)
				result[id] = true // Use map to avoid duplicates (e.g., 1111 = "1"×4 = "11"×2)
			}
		}
	}

	return result
}
