package _8A
import "fmt"

func Decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("started")
		f(s)
		fmt.Println("ended")
	}
}

