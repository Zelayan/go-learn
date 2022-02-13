package main

import (
	"fmt"
	"sort"
)

type ModelA struct {
	Id uint32
	Name string
}

type modelAs []ModelA

func (ps modelAs) Len() int {
	return len(ps)
}

func (ps modelAs) Less(i, j int) bool {
	return ps[i].Id < ps[j].Id
}

func (ps modelAs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	ps := modelAs{
		{1, "test"},
		{3, "test02"},
		{"lucy", "test03"},
	}
	sort.Sort(ps)
	fmt.Println(ps)
}
