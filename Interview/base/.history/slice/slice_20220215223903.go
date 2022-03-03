package slice

import "fmt"

func TwoArray() {
	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}

func NewArrayOrOldArray() {
	array := [4]int{10, 20, 30, 40}
	slice := array[0:2]
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	fmt.Printf("After array = %v\n", array)
}

func SliceCopy() {
	array := []int{10, 20, 30, 40}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n, slice)
}

func SliceCopyString() {
	slice := make([]byte, 3)
	n := copy(slice, "abcdef")
	fmt.Println(n, slice)
}

func RangeSliceIsValue() {
    slice := []int{10, 20, 30, 40}
    for index, value := range slice {
        fmt.Printf("value = %d , value-addr = %x , slice-addr = %x\n", value, &value, &slice[index])
    }

    // value = 10 , value-addr = c0000ac0c8 , slice-addr = c0000da060
    // value = 20 , value-addr = c0000ac0c8 , slice-addr = c0000da068
    // value = 30 , value-addr = c0000ac0c8 , slice-addr = c0000da070
    // value = 40 , value-addr = c0000ac0c8 , slice-addr = c0000da078
}

func SliceShaw o