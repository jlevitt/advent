package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type predicate func(string) bool

func hasSameLetter(times int) func(string) bool {
	return func(word string) bool {
		letters := map[rune]int{}

		for _, letter := range word {
			count, ok := letters[letter]

			if ok {
				letters[letter] = count + 1
			} else {
				letters[letter] = 1
			}
		}

		for _, count := range letters {
			if count == times {
				return true
			}
		}

		return false
	}
}

func countTrues(ids []string, pred predicate) int {
	count := 0

	for _, id := range ids {
		if pred(id) {
			count++
		}
	}

	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	ids := make([]string, 0)
	for scanner.Scan() {
		id := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		ids = append(ids, id)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	twos := countTrues(ids, hasSameLetter(2))
	threes := countTrues(ids, hasSameLetter(3))

	fmt.Printf("Twos: %v\nThrees: %v\nChecksum: %v\n", twos, threes, twos*threes)
}
