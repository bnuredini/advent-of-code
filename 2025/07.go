package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"maps"
)

func main() {
	f, err := os.Open("inputs/07.in")
	if err != nil {
		log.Fatalf("failed to read puzzle input: %v", err)
	}
	defer f.Close()

	answer1 := 0
	answer2 := 1

	scanner := bufio.NewScanner(f)
	beams := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()

		newBeams := make(map[int]int)
		maps.Copy(newBeams, beams)
		
		for colIdx, r := range line {
			if r == 'S' {
				newBeams[colIdx] = 1
				break
			} else if r == '^' && beams[colIdx] > 0 {
				answer1 += 1
				answer2 += beams[colIdx]

				if colIdx + 1 > 0 {
					newBeams[colIdx+1] += beams[colIdx]
				}

				if colIdx - 1 < len(line) {
					newBeams[colIdx-1] += beams[colIdx]
				}

				newBeams[colIdx] = 0
			}
		}

		beams = newBeams
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	fmt.Println(answer1)
	fmt.Println(answer2)
}

