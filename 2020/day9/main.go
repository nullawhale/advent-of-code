package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func fetchInput() (res []int) {
    data, err := ioutil.ReadFile("i")
    if err != nil {
	fmt.Printf("Err: %s", err)
    }
    str := strings.SplitN(string(data), "\n", 1000)
    for _, s := range str {
	num, _ := strconv.Atoi(strings.TrimSuffix(s, "\n"))
	res = append(res, num)
    }

    return res
}

func sum(slice []int, num int) bool {
    for i := 0; i < len(slice); i++ {
	for j := 0; j < len(slice); j++ {
	    if i != j {
		if slice[i] + slice[j] == num {
		    //fmt.Printf("%+v, %d = %d + %d\n", slice, num, slice[i], slice[j])
		    return true
		}
	    }
	}
    }
    return false
}

func star1() int {
    var idx = 0
    var data = fetchInput()

    for i := 25; i < len(data); i++ {
	if !sum(data[idx:idx+25], data[i]) {
	    return data[i]
	}
	idx++
    }

    return idx
}

func star2() int {
    var invalid = star1()
    var data = fetchInput()

    for i := 0; i < len(data); i++ {
	sum := data[i]
	for j := i + 1; j < len(data); j++ {
	    sum += data[j]

	    if sum == invalid {
		nums := data[i:j]
		sort.Ints(nums)
		fmt.Println(nums)
		min := nums[0]
		max := nums[len(nums)-1]
		return min + max
	    }

	    if sum > invalid {
		break
	    }
	}
    }

    return 0
}

func main() {
    //fmt.Printf("First star: %d\n", star1())
    fmt.Printf("Second star: %d\n", star2())
}
