package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func main() {
	f, err := os.Open("inputs/04.in")
	if err != nil {
		log.Fatalf("failed to read puzzle input: %v", err)
	}
	defer f.Close()

	answer1 := 0
	answer2 := 0
	mat := [][]bool{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		row := []bool{}
		for _, r := range line {
			if r == '@' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		mat = append(mat, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	answer1 = getRolls(mat, false)
	answer2 = getRolls(mat, true)

	fmt.Println(answer1)
	fmt.Println(answer2)
}

func getRolls(mat [][]bool, part2 bool) int {
	result := 0
	prevResult := 0
	running := true

	for running {
		idxMatToRemove := [][]int{}

		for rowIdx, row := range mat {
			idxRowToRemove := []int{}

			for colIdx, col := range row {
				if !col {
					continue
				}

				score := 0 

				score += getScore(mat, rowIdx - 1, colIdx - 1)
				score += getScore(mat, rowIdx - 1, colIdx)
				score += getScore(mat, rowIdx - 1, colIdx + 1)

				score += getScore(mat, rowIdx, colIdx - 1)
				score += getScore(mat, rowIdx, colIdx + 1)

				score += getScore(mat, rowIdx + 1, colIdx - 1)
				score += getScore(mat, rowIdx + 1, colIdx)
				score += getScore(mat, rowIdx + 1, colIdx + 1)

				
				if score < 4 {
					result += 1
					idxRowToRemove = append(idxRowToRemove, colIdx)
				}
			}

			idxMatToRemove = append(idxMatToRemove, idxRowToRemove)
		}

		if !part2 {
			// One pass-through the matrix is enough for part #1. Whereas, for 
			// part #2, multiple pass-through are required until no more rolls 
			// can be removed.
			running = false 
			break
		}

		for rowIdx, row := range idxMatToRemove {
			for _, col := range row {
				mat[rowIdx][col] = false
			}
		}

		if prevResult == result {
			running = false 
		}

		prevResult = result
	}

	return result
}

func getScore(mat [][]bool, rowIdx, colIdx int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}

	if rowIdx < 0 || rowIdx >= len(mat) {
		return 0 
	}

	if colIdx < 0 || colIdx >= len(mat[0]) {
		return 0
	}

	if mat[rowIdx][colIdx] {
		return 1
	}

	return 0
}

