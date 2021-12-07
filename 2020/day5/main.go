package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	//"regexp"
	//"strconv"
	"strings"
)

type Ticket struct {
	row int
	col int
	id int
}

func fetchInput() (res []string) {
	data, err := ioutil.ReadFile("i")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	res = strings.SplitN(string(data), "\n", 756)

	return res
}

func loadTickets() (tickets []Ticket){
	var data = fetchInput()

	//fmt.Printf("%+v", data)
	for _, ticket := range data {
		ticket = strings.TrimSuffix(ticket, "\n")
		ticketRow := ticket[:7]
		ticketCol := ticket[7:]

		F7min, F7max := 0, 127
		row := 0
		for _, r := range ticketRow {
			diff := ((F7max - F7min) / 2) + 1
			switch string(r) {
			case "F":
				row = F7min
				F7max -= diff
			case "B":
				row = F7max
				F7min += diff
			}
		}

		F3min, F3max := 0, 7
		col := 0
		for _, r := range ticketCol {
			diff := ((F3max - F3min) / 2) + 1
			switch string(r) {
			case "L":
				col = F3min
				F3max -= diff
			case "R":
				col = F3max
				F3min += diff
			}
		}

		id := row * 8 + col
		tickets = append(tickets, Ticket{ row: row, col: col, id: id })
	}

	return tickets
}

func star1() int {
	var maxId = 0
	var data = loadTickets()

	//fmt.Printf("%+v", data)
	for _, ticket := range data {
		if ticket.id > maxId {
		    maxId = ticket.id
		}
	}

	return maxId
}

func star2() int {
	var count = 0
	var data = loadTickets()

	//fmt.Printf("%+v", data)

	sort.Slice(data, func(i, j int) bool {
		return data[i].id < data[j].id
	})

	for i := 0; i < 127; i++ {
		for j := 0; j < 7; j++ {
			id := -127
			for _, ticket := range data {
				if ticket.row == i && ticket.col == j {
					id = ticket.id
					break
				}
			}
			if id != -127 {
				fmt.Printf("%d\t", id)
			} else {
				fmt.Printf(".\t")
			}
			id = -127
		}
		fmt.Println()
	}

	return count
}

func main() {
	//fmt.Printf("First star: %d\n", star1())
	fmt.Printf("Second star: %d\n", star2())
}
