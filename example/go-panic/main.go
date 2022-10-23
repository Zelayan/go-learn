package main

import "fmt"

/*func main() {
	recover()
	panic("it is panic") // not recover
}
*/

/*func main() {
	defer func() {
		recover()
	}()

	panic("it is panic") // recover
}
*/
/*
func main() {
	defer recover()
	panic("it is panic") // no recover
}*/

/*func main() {
	defer func() {
		defer func() {
			recover()
		}()
	}()

	panic("it is panic") // not recover
}
*/

/*func main() {
	defer doRecover()
	panic("it is panic") // recover
}

func doRecover() {
	defer recover()
}
*/

/*func main() {
	defer doRecover()
	panic("it is panic") // recover
}

func doRecover() {
	recover()
	fmt.Println("hello")
}*/

/*func main() {
	defer func() {
		defer func() {
			recover()
		}()
	}()

	panic("it is panic") // not recover
}*/

func main2() {
	go test()
	fmt.Println("main done")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recovery:%v", err)
		}
	}()
	panic("xx")
}
