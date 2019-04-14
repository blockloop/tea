package tea

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func newRequest() (*http.Request, *chi.Context) {
	req := httptest.NewRequest("GET", "/", nil)
	ctx := chi.NewRouteContext()
	ctx.RouteMethod = "GET"
	ctx.RoutePath = "/"
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	return req, ctx
}

func TestURLIntErrorsWhenParamDoesNotExist(t *testing.T) {
	r, _ := newRequest()
	i, err := URLInt(r, "id", "")
	assert.Equal(t, ErrNoURLParam, err)
	assert.EqualValues(t, 0, i)
}

func TestURLIntErrorsWhenNotAnInt(t *testing.T) {
	r, ctx := newRequest()
	ctx.URLParams.Add("id", "a")
	i, err := URLInt(r, "id", "")
	assert.Error(t, err)
	assert.EqualValues(t, 0, i)
}

func TestURLIntErrorsWhenValidationFails(t *testing.T) {
	r, ctx := newRequest()
	ctx.URLParams.Add("id", "10")
	_, err := URLInt(r, "id", "gt=10")
	assert.Error(t, err)
}

func TestURLIntReturnsTheValue(t *testing.T) {
	const expected = 10
	r, ctx := newRequest()
	ctx.URLParams.Add("id", fmt.Sprintf("%d", expected))
	i, err := URLInt(r, "id", "")
	assert.NoError(t, err)
	assert.EqualValues(t, expected, i)
}

func TestURLInt64ErrorsWhenParamDoesNotExist(t *testing.T) {
	r, _ := newRequest()
	i, err := URLInt64(r, "id", "")
	assert.Equal(t, ErrNoURLParam, err)
	assert.EqualValues(t, 0, i)
}

func TestURLInt64ErrorsWhenNotAnInt64(t *testing.T) {
	r, ctx := newRequest()
	ctx.URLParams.Add("id", "a")
	i, err := URLInt64(r, "id", "")
	assert.Error(t, err)
	assert.EqualValues(t, 0, i)
}

func TestURLInt64ErrorsWhenValidationFails(t *testing.T) {
	r, ctx := newRequest()
	ctx.URLParams.Add("id", "10")
	_, err := URLInt64(r, "id", "gt=10")
	assert.Error(t, err)
}

func TestURLInt64ReturnsTheValue(t *testing.T) {
	const expected = 10
	r, ctx := newRequest()
	ctx.URLParams.Add("id", fmt.Sprintf("%d", expected))
	i, err := URLInt64(r, "id", "")
	assert.NoError(t, err)
	assert.EqualValues(t, expected, i)
}

func TestURLUintErrorsWhenParamDoesNotExist(t *testing.T) {
	r, _ := newRequest()
	i, err := URLUint(r, "id", "")
	assert.Equal(t, ErrNoURLParam, err)
	assert.EqualValues(t, 0, i)
}

func TestURLUintErrorsWhenNotAUint(t *testing.T) {
	r, ctx := newRequest()
	ctx.URLParams.Add("id", "a")
	i, err := URLUint(r, "id", "")
	assert.Error(t, err)
	assert.EqualValues(t, 0, i)
}

func TestURLUintErrorsWhenValidationFails(t *testing.T) {
	r, ctx := newRequest()
	ctx.URLParams.Add("id", "10")
	_, err := URLUint(r, "id", "gt=10")
	assert.Error(t, err)
}

func TestURLUintReturnsTheValue(t *testing.T) {
	const expected = 10
	r, ctx := newRequest()
	ctx.URLParams.Add("id", fmt.Sprintf("%d", expected))
	i, err := URLUint(r, "id", "")
	assert.NoError(t, err)
	assert.EqualValues(t, expected, i)
}

func TestURLFloatErrorsWhenParamDoesNotExist(t *testing.T) {
	r, _ := newRequest()
	i, err := URLFloat(r, "id", "")
	assert.Equal(t, ErrNoURLParam, err)
	assert.EqualValues(t, 0, i)
}

func TestURLFloatErrorsWhenNotAFloat(t *testing.T) {
	r, ctx := newRequest()
	ctx.URLParams.Add("id", "a")
	i, err := URLFloat(r, "id", "")
	assert.Error(t, err)
	assert.EqualValues(t, 0, i)
}

func TestURLFloatErrorsWhenValidationFails(t *testing.T) {
	r, ctx := newRequest()
	ctx.URLParams.Add("id", "10")
	_, err := URLFloat(r, "id", "gt=10")
	assert.Error(t, err)
}

func TestURLFloatReturnsTheValue(t *testing.T) {
	const expected = 10
	r, ctx := newRequest()
	ctx.URLParams.Add("id", fmt.Sprintf("%d", expected))
	i, err := URLFloat(r, "id", "")
	assert.NoError(t, err)
	assert.EqualValues(t, expected, i)
}
