package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	chosenScores := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	loss := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	win := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}

	total := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		opponent := line[0]
		outcome := line[1]
		outcomeScore := 0

		var chosen string

		if outcome == "X" {
			chosen = loss[opponent]
		} else if outcome == "Y" {
			chosen = opponent
			outcomeScore = 3
		} else if outcome == "Z" {
			chosen = win[opponent]
			outcomeScore = 6
		}

		total += chosenScores[chosen] + outcomeScore
	}

	fmt.Println(total)
}
