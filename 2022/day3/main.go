package main

import (
	"fmt"
	"log"
	advent_of_code "nullawhale.com/aoc"
	"strings"
)

type Pair struct {
	left  string
	right string
}

func fetchInput(filename string) ([]string, error) {
	data, err := advent_of_code.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func star1(data []string) int {
	pairs := parsePairs(data)

	res := 0
	for _, pair := range pairs {
		var compartment rune
		for _, s := range pair.left {
			if stringContainsRune(pair.right, s) {
				compartment = s
			}
		}
		res += priority(compartment)
	}

	return res
}

func star2(data []string) int {
	pairs := parsePairs(data)

	res := 0

	for i := range pairs {
		var compartment rune
		if i%3 == 2 {
			for _, r := range data[i-2] {
				if stringContainsRune(data[i-1], r) && stringContainsRune(data[i], r) {
					compartment = r
				}
			}
		}
		res += priority(compartment)
	}

	return res
}

func stringContainsRune(line string, r rune) bool {
	return strings.Contains(line, string(r))
}

func priority(r rune) int {
	if 'a' <= r && r <= 'z' {
		return int(r - 'a' + 1)
	}
	if 'A' <= r && r <= 'Z' {
		return int(r - 'A' + 27)
	}
	return 0
}

func parsePairs(content []string) []Pair {
	var res []Pair
	for _, rucksack := range content {
		length := len(rucksack)
		res = append(res, Pair{
			left:  rucksack[0 : length/2],
			right: rucksack[length/2 : length],
		})
	}
	return res
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	data, err := fetchInput("i")
	if err != nil {
		log.Fatal(err)
	}

	star := star1(data)
	fmt.Printf("First star: %d\n", star)

	star = star2(data)
	fmt.Printf("Second star: %d\n", star)
}
