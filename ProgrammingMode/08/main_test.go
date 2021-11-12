package _8A

import (
	"fmt"
	"testing"
)

func test(s string) func(s string) {
	fmt.Println(s)
	return nil
}
func TestDecorator(t *testing.T) {
	f := test("12")
	//decorator := Decorator(f)
	f("1")

}
