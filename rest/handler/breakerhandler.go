package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/351423113/go-zero-extern/core/breaker"
	"github.com/351423113/go-zero-extern/core/logx"
	"github.com/351423113/go-zero-extern/core/stat"
	"github.com/351423113/go-zero-extern/rest/httpx"
	"github.com/351423113/go-zero-extern/rest/internal/security"
)

const breakerSeparator = "://"

// BreakerHandler returns a break circuit middleware.
func BreakerHandler(method, path string, metrics *stat.Metrics) func(http.Handler) http.Handler {
	brk := breaker.NewBreaker(breaker.WithName(strings.Join([]string{method, path}, breakerSeparator)))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			promise, err := brk.Allow()
			if err != nil {
				metrics.AddDrop()
				logx.Errorf("[http] dropped, %s - %s - %s",
					r.RequestURI, httpx.GetRemoteAddr(r), r.UserAgent())
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}

			cw := &security.WithCodeResponseWriter{Writer: w}
			defer func() {
				if cw.Code < http.StatusInternalServerError {
					promise.Accept()
				} else {
					promise.Reject(fmt.Sprintf("%d %s", cw.Code, http.StatusText(cw.Code)))
				}
			}()
			next.ServeHTTP(cw, r)
		})
	}
}
