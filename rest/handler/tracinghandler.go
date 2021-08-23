package handler

import (
	"net/http"

	"github.com/351423113/go-zero-extern/core/logx"
	"github.com/351423113/go-zero-extern/core/sysx"
	"github.com/351423113/go-zero-extern/core/trace"
)

// TracingHandler returns a middleware that traces the request.
func TracingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		carrier, err := trace.Extract(trace.HttpFormat, r.Header)
		// ErrInvalidCarrier means no trace id was set in http header
		if err != nil && err != trace.ErrInvalidCarrier {
			logx.Error(err)
		}

		ctx, span := trace.StartServerSpan(r.Context(), carrier, sysx.Hostname(), r.RequestURI)
		defer span.Finish()
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
