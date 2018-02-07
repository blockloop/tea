package tea

import (
	"net/http"

	"github.com/go-chi/render"
)

// Responder is the default responder used to write the messages back to the
// client. It uses render.DefaultResponder by default which will respond using
// the appropriate data type based on the Accept header. You can set it to
// render.JSON, render.XML, render.Data, etc. You can also create your own
// custom responder.
var Responder = render.DefaultResponder

// StatusHandlerFunc is a handler that returns a status code and a message body
type StatusHandlerFunc func(w http.ResponseWriter, r *http.Request) (int, interface{})

// Handler wraps a StatusHandlerFunc and returns a standard lib http.HandlerFunc
func Handler(h StatusHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, response := h(w, r)
		w.WriteHeader(status)
		if response == nil {
			return
		}

		Responder(w, r, response)
	}
}

// StatusError creates a response message consisting of the status code
// and the http.StatusText which applies to that code.
//
// This is useful within StatusHandlerFuncs to quickly break out of the
// normal flow of code
//
// Example:
//   func(w http.ResponseWriter, r *http.Request) (int, interface{}) {
//           u, p, ok := r.BasicAuth()
//           if !ok {
//                   return StatusError(StatusUnauthorized)
//	     }
//   }
func StatusError(status int) (int, *ErrorResponse) {
	return Error(status, http.StatusText(status))
}

// Error creates a response message consisting of the status code
// and the error string provided. The error string will be rendered back
// to the client. This is ideal for client errors where the client should
// be informed of the specific error message.
//
// This is useful within StatusHandlerFuncs to quickly break out of the
// normal flow of code. It also renders the status code in the body of
// the response which is often very helpful for clients.
//
// Example:
//   func CreateUser(w http.ResponseWriter, r *http.Request) (int, interface{}) {
//           // parse request
//           if req.Name == "" {
//                   return Error(400, "name is required")
//	     }
//   }
func Error(status int, err string) (int, *ErrorResponse) {
	return status, &ErrorResponse{
		Code:  status,
		Error: err,
	}
}

// ErrorResponse is a generic response object
type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
