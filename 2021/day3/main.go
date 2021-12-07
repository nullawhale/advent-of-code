package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func fetchInput() (res []string) {
	data, err := os.Open("i")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res
}

func star1() int64 {
	var gamma int64
	var epsilon int64
	var data = fetchInput()

	counts1 := make([]int, 12)
	for _, c := range data {
		for i, bit := range strings.Split(c, "") {
			//fmt.Println(i)
			if bit == "1" {
				counts1[i] += 1
			}
		}
	}

	var gammaString []string
	var epsilonString []string
	for _, count := range counts1 {
		if count > len(data)/2 {
			gammaString = append(gammaString, "1")
			epsilonString = append(epsilonString, "0")
		} else {
			gammaString = append(gammaString, "0")
			epsilonString = append(epsilonString, "1")
		}
	}
	gamma, _ = strconv.ParseInt(strings.Join(gammaString, ""), 2, 64)
	epsilon, _ = strconv.ParseInt(strings.Join(epsilonString, ""), 2, 64)

	return gamma * epsilon
}

func star2() int64 {
	var gamma int64
	var epsilon int64
	var data = fetchInput()

	gamma = CountRating(data, true)
	epsilon = CountRating(data, false)

	return gamma * epsilon
}

func CountRating(slice []string, keepMost bool) int64 {
	tempG := make([]string, len(slice))
	copy(tempG, slice)

	var cmpString string
	for i := range tempG {
		counts := 0.0
		for _, bit := range tempG {
			if string(bit[i]) == "1" {
				counts += 1
			}
		}

		l := math.Ceil(float64(len(tempG)) / 2)
		if counts == l {
			if keepMost {
				cmpString += "1"
			} else {
				cmpString += "0"
			}
		} else if counts > l {
			if keepMost {
				cmpString += "1"
			} else {
				cmpString += "0"
			}
		} else {
			if keepMost {
				cmpString += "0"
			} else {
				cmpString += "1"
			}
		}

		for j := 0; j < len(tempG); {
			if !strings.HasPrefix(tempG[j], cmpString) {
				tempG = remove(tempG, j)
			} else {
				j++
			}
		}

		if len(tempG) == 1 {
			res, _ := strconv.ParseInt(tempG[0], 2, 64)
			return res
		}
	}
	return 0
}

func remove(slice []string, i int) []string {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	fmt.Printf("First star: %d\n", star1())
	fmt.Printf("Second star: %d\n", star2())
}
