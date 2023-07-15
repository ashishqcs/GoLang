package factorial

import "testing"

func TestShouldReturn1WhenInputIs1(t *testing.T) {
	expected := 1
	actual := factorial(1)

	if expected != actual {
		t.Errorf("Test failed: Wanted %d but Got %d", expected, actual)
	}
}

func TestShouldReturn6WhenInputIs3(t *testing.T) {
	expected := 6
	actual := factorial(3)

	if expected != actual {
		t.Errorf("Test failed: Wanted %d but Got %d", expected, actual)
	}
}

func TestShouldReturn24WhenInputIs4(t *testing.T) {
	expected := 24
	actual := factorial(4)

	if expected != actual {
		t.Errorf("Test failed: Wanted %d but Got %d", expected, actual)
	}
}

func TestShouldReturn1WhenInputIs0(t *testing.T) {
	expected := 1
	actual := factorial(0)

	if expected != actual {
		t.Errorf("Test failed: Wanted %d but Got %d", expected, actual)
	}
}

func TestShouldReturn40320WhenInputIs6(t *testing.T) {
	expected := 40320
	actual := factorial(8)

	if expected != actual {
		t.Errorf("Test failed: Wanted %d but Got %d", expected, actual)
	}
}
