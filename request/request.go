package request

import (
	"context"
	"encoding/json"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

// Validate is the Validator used to validate structs within Body
var Validate Validator = validator.New()

// Validator is a struct validator used within Body
type Validator interface {
	StructCtx(context.Context, interface{}) error
}

// Body parses a JSON request body and validates it's contents
func Body(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return Validate.StructCtx(r.Context(), v)
}
