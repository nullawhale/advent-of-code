package main

import (
	"fmt"
	"log"
	advent_of_code "nullawhale.com/aoc"
	"sort"
	"strconv"
)

func fetchInput(filename string) ([]string, error) {
	data, err := advent_of_code.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func star1(data []string) (int, error) {
	elves, err := rowSums(data)
	if err != nil {
		return 0, err
	}

	max := 0
	for _, elfCalories := range elves {
		if elfCalories > max {
			max = elfCalories
		}
	}

	return max, nil
}

func star2(data []string) (int, error) {
	elves, err := rowSums(data)
	if err != nil {
		return 0, err
	}

	values := make([]int, 0, len(elves))
	for _, elf := range elves {
		values = append(values, elf)
	}

	sort.Ints(values)

	return values[len(values)-1] + values[len(values)-2] + values[len(values)-3], nil
}

func rowSums(strings []string) (map[int]int, error) {
	res := map[int]int{}
	index := 1
	for _, s := range strings {
		if s == "" {
			index++
		} else {
			calorie, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			res[index] += calorie
		}
	}
	return res, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	data, err := fetchInput("i")
	if err != nil {
		log.Fatal(err)
	}

	star, err := star1(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("First star: %d\n", star)

	star, err = star2(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Second star: %d\n", star)
}
