package tea

import (
	"testing"
)

func TestErrorReturnsSameCode(t *testing.T) {
	expected := 400
	actual, _ := Error(expected, "")
	if expected != actual {
		t.Fatalf("expected '%d', got '%d'", expected, actual)
	}
}

func TestErrorReturnsErrorWithSameCode(t *testing.T) {
	expected := "bad news"
	_, err := Error(400, expected)
	actual := err.Error()
	if expected != actual {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}
