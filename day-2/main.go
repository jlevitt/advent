package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func hasAnyLetterTwice(word string) bool {
	letters := map[rune]int{}

	for _, letter := range word {
		count, ok := letters[letter]

		if ok {
			letters[letter] = count + 1
		} else {
			letters[letter] = 1
		}

		if letters[letter] == 2 {
			return true
		}
	}

	return false
}

func countTwos(ids []string) int {
	count := 0

	for _, id := range ids {
		if hasAnyLetterTwice(id) {
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

	twos := countTwos(ids)

	fmt.Printf("Final id: %v\n", twos)
}
