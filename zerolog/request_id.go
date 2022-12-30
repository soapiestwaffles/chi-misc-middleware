package zerolog

import (
	"net/http"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ZerologRequestID configures a ZeroLog sub-logger with requestID in every log entry
// related to an HTTP request. Requires the Chi requestID middleware.
func ZerologRequestID() func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := log.With().
				Dict("httpRequest", zerolog.Dict().
					Str("requestID", chimiddleware.GetReqID(r.Context()))).
				Logger().
				WithContext(r.Context())

			h.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
	return f
}
