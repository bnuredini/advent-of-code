package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/02.in")
	if err != nil {
		log.Fatalf("faliled to open the input file: %v", err)
	}
	defer f.Close()

	answer1 := 0
	answer2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		intervals := strings.Split(line, ",")
		if len(intervals) == 0 {
			continue
		}

		for _, interval := range intervals {
			parts := strings.Split(interval, "-")
			if len(parts) == 0 {
				log.Fatal("invalid puzzle input")
			}

			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			for i := start; i <= end; i++ {
				str := strconv.Itoa(i)
				numDigits := len(str)

				if numDigits == 1 {
					continue
				}

				if isInvalid(str, false) {
					answer1 += i
				}

				if isInvalid(str, true) {
					answer2 += i
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	fmt.Println(answer1)
	fmt.Println(answer2)
}

func isInvalid(s string, part2 bool) bool {
	numDigits := len(s)

	start := 0
	if part2 {
		start = numDigits
	} else {
		start = 2
	}

	for i := start; i >= 2; i-- {
		if numDigits%i == 0 {
			numChunks := i
			chunkLength := numDigits / i
			shouldSkip := false

			for j := 0; j < numChunks-1; j++ {
				lastChunk := s[numDigits-chunkLength:]

				if s[j*chunkLength:(j+1)*chunkLength] != lastChunk {
					shouldSkip = true
					break
				}
			}

			if !shouldSkip {
				return true
			}
		}
	}

	return false
}
