package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/01.in")
	if err != nil {
		log.Fatalf("faliled to open the input file: %v", err)
	}
	defer f.Close()

	position := 50
	answer1 := 0
	answer2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 2 {
			continue
		}

		direction := line[0]
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		if strings.ToUpper(string(direction)) == "L" {
			oldPosition := position 
			position = position - amount

			if position == 0 {
				answer1 += 1
				answer2 += 1
			} else {
				for position < 0 {
					position += 100
					answer2 += 1
				}

				// Edge case: if the loop above ends at zero, it doesn't count the 
				// last click since it just just counts the number of wraps. 
				if position == 0 {
					answer1 += 1
					answer2 += 1
				}

				// Edge case: going from zero to a negative value shouldn't cause a
				// click. But, the loop above does when it counts the number of 
				// wraps.
				if oldPosition == 0 && position > 0 {
					answer2 -= 1
				}
			}
		} else if strings.ToUpper(string(direction)) == "R" {
			position = position + amount
			for position >= 100 {
				position -= 100
				answer2 += 1
			}

			if position == 0 {
				answer1 += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}
	
	fmt.Println(answer1)
	fmt.Println(answer2)
}
