package tea

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatusErrorReturnsStatusCodeWithStatusText(t *testing.T) {
	const expected = http.StatusConflict
	code, err := StatusError(expected)
	assert.Error(t, err)
	assert.EqualValues(t, expected, code)
	assert.Equal(t, http.StatusText(expected), err.Error())
}

func TestErrorfFormatsError(t *testing.T) {
	f := "name %q, age: %d"
	args := []interface{}{"brett", "100"}
	expected := fmt.Sprintf(f, args...)

	_, err := Errorf(400, f, args...)
	require.Error(t, err)
	assert.Equal(t, expected, err.Error())
}

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
	require.Error(t, err)
	actual := err.Error()
	if expected != actual {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestNotFoundUses404Code(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
        NotFound(w, r)
	w.Flush()
	assert.Equal(t, 404, w.Code)
}

func TestNotFoundUses405Code(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	MethodNotAllowed(w, r)
	w.Flush()
	assert.Equal(t, 405, w.Code)
}
