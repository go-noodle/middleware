package middleware

import (
	"log"
	"net/http"

	"github.com/go-noodle/noodle"
)

// Default is a convenience function creating new noodle.Chain with Logger, Recover and LocalStore middlewares
func Default(mws ...noodle.Middleware) noodle.Chain {
	return noodle.New(RealIP, Logger, Recover(func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("%+v", err)
		http.Error(w, err.Error(), 500)
	}), LocalStore).Use(mws...)
}
