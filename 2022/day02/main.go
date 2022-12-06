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
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	draws := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	wins := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	total := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		opponent := line[0]
		chosen := line[1]

		outcome := 0
		if draws[opponent] == chosen {
			outcome = 3
		} else if wins[opponent] == chosen {
			outcome = 6
		}

		fmt.Println(line, outcome)

		total += chosenScores[chosen] + outcome
	}

	fmt.Println(total)
}
