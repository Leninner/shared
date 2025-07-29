package middleware

import (
	"fmt"
	"net/http"

	"github.com/leninner/shared/exception"
)

func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			pv := recover()
			if pv != nil {

				w.Header().Set("Connection", "close")

				exception.ServerErrorResponse(w, r, fmt.Errorf("%v", pv))
			}
		}()

		next.ServeHTTP(w, r)
	})
} 