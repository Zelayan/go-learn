package list

import "container/list"

func List() {


	l := list.New()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// 遍历
for e := l.Front(); e != nil; e = e.Next() {
	fmt.Printf("%v\n", e.Value)
}

}