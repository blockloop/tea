package tea

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync/atomic"
	"testing"
)

func TestHandlerWritesStatusCode(t *testing.T) {
	expected := http.StatusBadRequest
	h := Handler(func(http.ResponseWriter, *http.Request) (int, interface{}) {
		return expected, nil
	})

	w := httptest.NewRecorder()
	h.ServeHTTP(w, nil)
	w.Flush()

	if expected != w.Code {
		t.Fatalf("expected %d, got %d", expected, w.Code)
	}
}

func TestHandlerWritesWithResponder(t *testing.T) {
	const yes = 1
	var called int32

	body := struct{ Name string }{"Brett"}

	Responder = func(w http.ResponseWriter, r *http.Request, v interface{}) {
		atomic.StoreInt32(&called, yes)
		if !reflect.DeepEqual(v, body) {
			t.Fatalf("expected body to be %+v, but was %+v", body, v)
		}
	}

	h := Handler(func(http.ResponseWriter, *http.Request) (int, interface{}) {
		return 200, body
	})

	w := httptest.NewRecorder()
	h.ServeHTTP(w, nil)
	w.Flush()

	if atomic.LoadInt32(&called) != yes {
		t.Fatalf("expected Responder to be called")
	}
}

func TestHandlerDoesNotWriteIfBodyIsNil(t *testing.T) {
	const yes = 1
	var called int32

	Responder = func(w http.ResponseWriter, r *http.Request, v interface{}) {
		atomic.StoreInt32(&called, yes)
	}

	h := Handler(func(http.ResponseWriter, *http.Request) (int, interface{}) {
		return 200, nil
	})

	w := httptest.NewRecorder()
	h.ServeHTTP(w, nil)
	w.Flush()

	if atomic.LoadInt32(&called) == yes {
		t.Fatalf("expected Responder to NOT be called")
	}
}
