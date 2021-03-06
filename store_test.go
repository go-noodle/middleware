package middleware_test

import (
	"github.com/go-noodle/noodle"
	mw "github.com/go-noodle/middleware"
	"gopkg.in/tylerb/is.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStore(t *testing.T) {
	is := is.New(t)
	n := noodle.New(mw.LocalStore).Then(
		func(w http.ResponseWriter, r *http.Request) {
			s := mw.GetStore(r)
			is.NotNil(s)
		},
	)
	r, _ := http.NewRequest("GET", "http://localhost", nil)
	n(httptest.NewRecorder(), r)
}
