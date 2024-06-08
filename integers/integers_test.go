package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	displayHelper(t, sum, expected)
}

func displayHelper(t testing.TB, got, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("Expected :  %d, but got : %d", got, expected)
	}
}
