package zerolog

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

// ZerologOtelTraceID adds traceID to the Zerolog logging context
func ZerologOtelTraceID() func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			span := trace.SpanFromContext(r.Context())

			// set up traceID in logs
			ctx = log.With().
				Str("traceID", span.SpanContext().TraceID().String()).
				Logger().
				WithContext(ctx)

			h.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
	return f
}
