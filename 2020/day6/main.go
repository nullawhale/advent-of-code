package main

import (
	"fmt"
	"io/ioutil"
	//"regexp"
	//"strconv"
	"strings"
)

type Answers struct {
	groups []Group
}

type Group struct {
	group []string
}

func unique(slice []string) (res []string) {
	keys := make(map[string]bool)
	for _, entry := range slice {
		if !keys[entry] {
			res = append(res, entry)
			keys[entry] = true
		}
	}
	return res
}

func intersection(s1, s2 []string) (res []string) {
	keys := make(map[string]bool)
	for _, e := range s1 {
		keys[e] = true
	}
	for _, e := range s2 {
		if keys[e] {
			res = append(res, e)
		}
	}

	return res
}

func fetchInput() (ans []Answers) {
	data, err := ioutil.ReadFile("i")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	str := strings.Split(string(data), "\n\n")
	for _, s := range str {
		s = strings.TrimSuffix(s, "\n")
		groups := strings.Split(s, "\n")
		var a Answers
		for _, g := range groups {
			gt := strings.Split(g, "")
			a.groups = append(a.groups, Group{group: gt})
		}
		ans = append(ans, a)
	}

	return ans
}

func star1() int {
	var count = 0
	var data = fetchInput()

	//fmt.Printf("%+v", data)

	for _, i := range data {
		var g []string
		for _, j := range i.groups {
			g = append(g, j.group...)
		}
		count += len(unique(g))
	}

	return count
}

func star2() int {
	var count = 0
	var data = fetchInput()

	//fmt.Printf("%+v", data)

	for _, i := range data {
		//fmt.Printf("%d\n", len(i.groups))
		if len(i.groups) > 1 {
			var interSlice []string
			interSlice = intersection(i.groups[0].group, i.groups[1].group)
			for j := 2; j < len(i.groups); j++ {
				interSlice = intersection(interSlice, i.groups[j].group)
			}
			count += len(interSlice)
		} else {
			count += len(i.groups[0].group)
		}

	}

	return count
}

func main() {
	fmt.Printf("First star: %d\n", star1())
	fmt.Printf("Second star: %d\n", star2())
}
