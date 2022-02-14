package iteration

func Repeat(character string) string {

	var repeadted string 
	for i := 0; i < 5; i++ {
		repeadted = repeadted + character
	}
	return repeadted
}