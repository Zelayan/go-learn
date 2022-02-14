package iteration

func Repeat(character string) string {

	var repeadted string 
	x := 5
	for i := 0; i < x; i++ {
		repeadted = repeadted + character
	}
	return repeadted
}