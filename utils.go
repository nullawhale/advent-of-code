package advent_of_code

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	stringContent := strings.TrimSpace(string(content))
	return strings.Split(stringContent, "\n"), nil
}

func ReadFileAsIs(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	input := string(content)
	input = strings.Trim(input, "\n")

	inputSplit := strings.Split(input, "\n\n")
	return inputSplit, nil
}

func StringSliceToIntSlice(stringSlice []string) ([]int, error) {
	var intSlice []int

	for _, v := range stringSlice {
		res, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		intSlice = append(intSlice, res)
	}

	return intSlice, nil
}
