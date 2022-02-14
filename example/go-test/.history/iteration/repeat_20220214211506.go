package iteration

const repeatCount = 5

func Repeat(character string, count int) string {

	var repeated string 
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}