package arrays

import (
	"reflect"
	"testing"

	"github.com/influxdata/influxdb/kit/check"
)


func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if want != got {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if want != got {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{2, 3})
	want := []int{6, 5}
	// Go语言中不对切片进行等号运算符
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumTails(t *testing.T) {

	checkNums := func 	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
()  {
		
	}
	got := SumTails([]int{1, 2}, []int{2, 3})
	want := []int{2, 3}


	t.Run("empty slice", func(t *testing.T) {
		got := SumTails([]int{}, []int{2, 3})
		want := []int{0, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	
	})
}