package generic

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println("1 + 2 =", Add[int64](1, 2))
}
func TestDel(t *testing.T) {
	t.Log(Add[int64](1, 2))
}
func TestTest(t *testing.T) {
	t.Log("xxx")
}
