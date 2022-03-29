package generic

func Add2[T int | int64](a, b T) T {
	return a + b
}
func Sub[T int | int64](a, b T) T {
	return a - b
}
func Mulit[T int | int64](a, b T) T {
	return a * b
}
