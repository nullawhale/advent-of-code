package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Instructions struct {
	inst     	string
	arg 		int
}

func fetchInput() (instructions []Instructions) {
	data, err := ioutil.ReadFile("i")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	str := strings.Split(string(data), "\n")

	for _, i := range str {
		//fmt.Printf("%s\n", i)
		inst := strings.Split(i, " ")[0]
		arg, _ := strconv.Atoi(strings.Split(i, " ")[1])
		//fmt.Printf("%s, %d\n", inst, arg)
		instructions = append(instructions, Instructions{inst, arg})
	}

	return instructions
}

func star1() int {
	var acc = 0
	var data = fetchInput()
	var indexes = make(map[int]bool)

	//fmt.Printf("%+v", data)

	var i = 0
	for !indexes[i] {
		//fmt.Printf("%d: %s, %d\n", i, data[i].inst, data[i].arg)
		indexes[i] = true
		inst := data[i].inst
		arg := data[i].arg

		switch inst {
		case "acc":
			acc += arg
			i++
		case "jmp":
			i += arg
		case "nop":
			i++
		}
	}

	return acc
}

func star2() int {
	var data = fetchInput()

	//fmt.Printf("%+v", data)

	for j := 0; j < len(data); j++ {
		//fmt.Printf("%+v", data)

		var acc = 0
		var indexes = make(map[int]bool)
		var i = 0
		for !indexes[i] {
			if i >= len(data) {
				return acc
			}
			//fmt.Printf("%d: %s, %d\n", i, data[i].inst, data[i].arg)
			indexes[i] = true
			inst := data[i].inst
			arg := data[i].arg

			if j == i {
				if inst == "jmp" {
					inst = "nop"
				} else if inst == "nop" {
					inst = "jmp"
				}
			}

			switch inst {
			case "acc":
				acc += arg
				i++
			case "jmp":
				i += arg
			case "nop":
				i++
			}
		}
	}

	return 0
}

func main() {
	fmt.Printf("First star: %d\n", star1())
	fmt.Printf("Second star: %d\n", star2())

	//fetchInput()
}
