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

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line)-3; i++ {
			duplicate := false
			for j := 0; j < 14; j++ {
				if strings.Count(line[i:i+14], string(line[i+j])) > 1 {
					duplicate = true
					break
				}
			}

			if duplicate == false {
				fmt.Println("found at: ", i+14)
				break
			}
		}

	}
}
