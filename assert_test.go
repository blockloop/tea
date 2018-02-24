package tea

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected '%+v', got '%+v'", expected, actual)
	}
}

func assertNotEqual(t *testing.T, expected, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected '%+v' to NOT equal '%+v'", actual, expected)
	}
}
