package main

import (
	"runtime"
	"testing"
)

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024) // 1M
		ans = append(ans, f(origin))
		// GC
		runtime.GC()
	}
	printMem(t)
	_ = ans
}

// 100/14m
func TestLastCharsBySlice(t *testing.T) {
	testLastChars(t, lastNumsBySlice)
}
// 1m
func TestLastCharsByCopy(t *testing.T)  {
	testLastChars(t, lastNumsByCopy)
}