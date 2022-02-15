package slice

import "testing"

func TestTwoArray(t *testing.T) {
	TwoArray()
}

func TestNewArrayOrOldArray(t *testing.T) {
	NewArrayOrOldArray()
	// Before slice = [10 20], Pointer = 0xc0000c4060, len = 2, cap = 4
	// Before newSlice = [10 20 50], Pointer = 0xc0000c4078, len = 3, cap = 4
	// After slice = [10 30], Pointer = 0xc0000c4060, len = 2, cap = 4
	// After newSlice = [10 30 50], Pointer = 0xc0000c4078, len = 3, cap = 4
	// After array = [10 30 50 40]
}

func TestSliceCopy(t *testing.T) {
	SliceCopy()
}

func TestSliceCopyString(t *testing.T) {
	SliceCopyString()
}

//由于 Value 是值拷贝的，并非引用传递，所以直接改 Value 是达不到更改原切片值的目的的，需要通过 &slice[index] 获取真实的地址。
func TestRangeSliceIsAnValue(t *testing.T) {
	RangeSliceIsValue()
}

//浅ka
func TestSliceShallowCopy(t *testing.T) {
	SliceShallowCopy()
}

//深拷贝：源切片和目的切片各自都有彼此独立的底层数组空间，各自的修改，彼此不受影响
func TestSliceDeepCopy(t *testing.T) {
	SliceDeepCopy()
}

