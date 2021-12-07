package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type PassPolicy struct {
	minC int
	maxC int
	char string
	pass string
}

func fetchInput() (passPolicies []PassPolicy) {
	data, err := ioutil.ReadFile("i")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	str := strings.SplitN(string(data), "\n", 1000)
	for _, s := range str {
		s = strings.TrimSuffix(s, "\n")
		//fmt.Printf("'%s'\n", s)

		r, _ := regexp.Compile(`(\d*)-(\d*)\s(\w):\s(\w*)`)
		matches := r.FindStringSubmatch(s)
		matches = matches[1:]
		min, _ := strconv.Atoi(matches[0])
		max, _ := strconv.Atoi(matches[1])

		passPolicies = append(
			passPolicies,
			PassPolicy{min, max, matches[2], matches[3]})
	}

	return passPolicies
}

func countSubstings(s string, sub string) int {
	r := regexp.MustCompile(sub)
	matches := r.FindAllStringIndex(s, -1)
	return len(matches)
}

func star1() int {
	var count = 0
	var data = fetchInput()

	//fmt.Printf("%+v", data)
	for _, p := range data {
		num := countSubstings(p.pass, p.char)
		if num >= p.minC && num <= p.maxC {
			count++
		}
	}

	return count
}

func star2() int {
	var count = 0
	var data = fetchInput()

	//fmt.Printf("%+v", data)
	for _, p := range data {
		//fmt.Printf("%d-%d %s: %s\n", p.minC, p.maxC, p.char, p.pass)
		if (string(p.pass[p.minC-1]) == p.char && string(p.pass[p.maxC-1]) != p.char) ||
			(string(p.pass[p.minC-1]) != p.char && string(p.pass[p.maxC-1]) == p.char) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Printf("First star: %d\n", star1())
	fmt.Printf("Second star: %d\n", star2())
}
