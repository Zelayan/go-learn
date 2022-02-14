package integers

import "testing"

func TestAdder(t *testing.T) {
	
	sum := Add(2, 2)
	excepted := 4

	if sum != excepted {
		t.Errorf("excepted %d but got %d", excepted, sum)
	}
}

func Example() {
	sum := Add()
	//Output:
	
}