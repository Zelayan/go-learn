package iteration

const repeatCount = 5

func Repeat(character string) string {

	var repeadted string 
	for i := 0; i < repeatCount; i++ {
		repeadted = repeadted + character
	}
	return repeadted
}