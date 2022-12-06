package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	calories := make([]int, 0)
	count := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			calories = append(calories, count)
			count = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		count += num
	}

	calories = append(calories, count)

	slices.Sort(calories)
	fmt.Println(calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3])
}
