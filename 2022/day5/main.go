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
	data, err := advent_of_code.ReadFileAsIs(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func star1(data []string) int {
	stacks := convertData(data[0])
	fmt.Printf("%v\n", stacks)
	/*crate_stacks := crateStacksFromStr(inputSplit[0])
	instr_strs := strings.Split(strings.Trim(inputSplit[1], "\n"), "\n")
	pairs, err := convertData(data)
	if err != nil {
		log.Fatal(err)
	}*/

	res := 0

	return res
}

func star2(data []string) int {
	/*pairs, err := convertData(data)
	if err != nil {
		log.Fatal(err)
	}*/

	res := 0

	return res
}

func convertData(stacks string) [][]string {
	lines := strings.Split(stacks, "\n")
	stackCount := (len(lines[0]) - 2) / 3

	result := make([][]string, stackCount)

	for i := 0; i < stackCount; i++ {
		for j := len(lines) - 2; j >= 0; j-- {
			if lines[j][i*4+1] == ' ' {
				break
			}
			result[i] = append(result[i], string(lines[j][i*4+1]))
		}
	}

	return result
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	data, err := fetchInput("i_test")
	if err != nil {
		log.Fatal(err)
	}

	star := star1(data)
	fmt.Printf("First star: %d\n", star)

	star = star2(data)
	fmt.Printf("Second star: %d\n", star)
}
