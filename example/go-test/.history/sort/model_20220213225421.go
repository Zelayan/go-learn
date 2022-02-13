package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Id uint32
	Name string
}

type persons []Person

func (ps persons) Len() int {
	return len(ps)
}

func (ps persons) Less(i, j int) bool {
	return ps[i].Id < ps[j].Id
}

func (ps persons) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	ps := persons{
		{"larry", 19},
		{"jackey", 18},
		{"lucy", 20},
	}
	sort.Sort(ps)
	fmt.Println(ps)
}
