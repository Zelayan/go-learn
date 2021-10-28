package main

import "fmt"

type GraphMap map[string][]string

func main() {

	var graphMap = make(GraphMap, 0)

	graphMap["you"] = []string{"jack", "rose", "tt"}
	graphMap["jack"] = []string{"less", "123"}
	graphMap["rose"] = []string{"mm", "aa"}
	graphMap["tt"] = []string{"tom"}

    // 模拟队列
	searchQueue := graphMap["you"]

	for {
		if len(searchQueue) > 0 {
			var person string
			person, searchQueue = searchQueue[0], searchQueue[1:]
			if personIsTom(person) {
				fmt.Printf("%s is the man\n", person)
				break
			} else {
				// 依次入队
				searchQueue = append(searchQueue, graphMap[person]...)
			}
		} else {
			fmt.Println("not found")
			break
		}

	}


}

func personIsTom(person string) bool {
	return person == "tom"
}
