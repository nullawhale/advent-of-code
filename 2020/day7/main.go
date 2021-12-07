package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Bag struct {
	color     string
	innerBags []InnerBag
}

type InnerBag struct {
	count int
	color string
}

func fetchInput() (bags []Bag) {
	data, err := ioutil.ReadFile("i")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	str := strings.Split(string(data), "\n")

	var bag Bag
	for _, s := range str {
		s = strings.TrimSuffix(s, "\n")
		color := strings.Split(s, " bags contain ")
		if !strings.HasSuffix(color[1], "no other bags.") {
			var innerBags []InnerBag
			innerBag := strings.Split(color[1], ", ")
			for _, ib := range innerBag {
				ib = strings.TrimSuffix(ib, ".")
				ib = strings.TrimSuffix(ib, " bag")
				ib = strings.TrimSuffix(ib, " bags")
				ibCount, _ := strconv.Atoi(strings.SplitN(ib, " ", 2)[0])
				ibColor := strings.SplitN(ib, " ", 2)[1]
				innerBags = append(innerBags, InnerBag{
					count: ibCount,
					color: ibColor,
				})
			}

			bag = Bag{
				color:     color[0],
				innerBags: innerBags,
			}
		} else {
			bag = Bag{
				color:     color[0],
				innerBags: nil,
			}
		}

		bags = append(bags, bag)
	}

	/*for _, bbb := range bags {
		fmt.Printf("%s\n", bbb.color)
		for _, ibbb := range bbb.innerBags {
			fmt.Printf("\t%s (%d)\n", ibbb.color, ibbb.count)
		}
	}*/

	return bags
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

func recursion(bags []Bag, color string, res *[]string) {
	var containing []string

	for _, bag := range bags {
		for _, innerBag := range bag.innerBags {
			if innerBag.color == color {
				containing = append(containing, bag.color)
			}
		}
	}

	for _, key := range containing {
		recursion(bags, key, res)
		*res = append(*res, key)
	}
}

func star1() int {
	var res []string
	var data = fetchInput()

	//fmt.Printf("%+v", data)
	recursion(data, "shiny gold", &res)
	res = unique(res)

	return len(res)
}

func star2(color string) int {
	var count = 0
	var data = fetchInput()

	//fmt.Printf("%+v", data)

	for _, bag := range data {
		for _, innerBag := range bag.innerBags {
			if bag.color == color {
				count += innerBag.count + innerBag.count * star2(innerBag.color)
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("First star: %d\n", star1())
	fmt.Printf("Second star: %d\n", star2("shiny gold"))

	//fetchInput()
}
