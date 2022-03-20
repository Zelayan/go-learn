package generic

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println("1 + 2 =", Add[int64](1, 2))
}