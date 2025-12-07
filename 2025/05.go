package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/05.in")
	if err != nil {
		log.Fatalf("failed to read puzzle input: %v", err)
	}
	defer f.Close()

	answer1 := 0
	answer2 := 0

	parsingIDs := false
	intervals := [][]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if !parsingIDs {
				parsingIDs = true
				continue
			} else {
				break
			}
		}

		if parsingIDs {
			id, _ := strconv.Atoi(line)
			for _, interval := range intervals {
				if id >= interval[0] && id <= interval[1] {
					answer1 += 1
					break
				}
			}
		} else {
			parts := strings.Split(line, "-")
			if len(parts) == 0 {
				log.Fatal("invalid puzzle input")
			}

			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			intervals = append(intervals, []int{start, end})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start := intervals[0][0]
	end := intervals[0][1]

	for _, interval := range intervals[1:] {
		newStart := interval[0]
		newEnd := interval[1]

		if newStart > end {
			answer2 += (end - start) + 1

			start = newStart
			end = newEnd
		} else if newEnd > end {
			end = newEnd
		}
	}

	answer2 += (end - start) + 1

	fmt.Println(answer1)
	fmt.Println(answer2)
}
