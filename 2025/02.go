package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strings"
	"strconv"
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
			startRaw := parts[0]
			endRaw := parts[1]

			start, _ := strconv.Atoi(startRaw)
			end, _ := strconv.Atoi(endRaw)

			for i := start; i <= end; i++ {
				str := strconv.Itoa(i)

				numDigits := len(str)
				if numDigits == 1 {
					continue
				} 

				if numDigits == 2 {
					if isJustOneRepeatedDigit(str) {
						answer1 += i
						answer2 += i
					}

					continue
				}

				if numDigits % 2 == 0 {
					divisors := []int{}

					for j := (numDigits/2); j >= 2; j-- {
						if numDigits % j == 0 {
							divisors = append(divisors, j)
						}
					}

					for _, divisor := range divisors {
						chunkLength := divisor
						numChunks := numDigits / chunkLength
						shouldSkip := false

						for chunkIdx := 0; chunkIdx < numChunks - 1; chunkIdx++ {
							chunkRaw := str[(chunkIdx)*chunkLength:(chunkIdx+1)*chunkLength]
							nextChunkRaw := str[(chunkIdx+1)*chunkLength:(chunkIdx+2)*chunkLength]

							chunk, _ := strconv.Atoi(chunkRaw)
							nextChunk, _ := strconv.Atoi(nextChunkRaw)

							if chunk - nextChunk != 0 {
								shouldSkip = true
								break
							}
						}

						if !shouldSkip {
							if numChunks == 2 {
								answer1 += i
							}

							answer2 += i
							break
						}
					}
				} else {
					if isJustOneRepeatedDigit(str) { 
						answer2 += i
					}

					divisors := []int{}
					for j := (numDigits/2) - 1; j >= 3; j-- {
						if numDigits % j == 0 {
							divisors = append(divisors, j)
						}
					}

					if len(divisors) == 0 {
						continue
					}

					for _, divisor := range divisors {
						chunkLength := divisor
						numChunks := numDigits / chunkLength
						noMatch := false

						for chunkIdx := 0; chunkIdx < numChunks - 1; chunkIdx++ {
							chunkRaw := str[(chunkIdx)*chunkLength:(chunkIdx+1)*chunkLength]
							nextChunkRaw := str[(chunkIdx+1)*chunkLength:(chunkIdx+2)*chunkLength]

							chunk, _ := strconv.Atoi(chunkRaw)
							nextChunk, _ := strconv.Atoi(nextChunkRaw)

							if chunk - nextChunk != 0 {
								noMatch = true
								break
							}
						}

						if !noMatch {
							answer2 += i
						}
					}
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

func isJustOneRepeatedDigit(s string) bool {
	for i := 0; i < len(s) - 1; i++ {
		if s[i] != s[i+1] {
			return false
		}
	}

	return true
}
