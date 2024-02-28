package middlewares

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackjohn7/smllnk/app"
	"github.com/urfave/negroni"
)

type Logger struct {
	server  *http.ServeMux
	slogger *slog.Logger
}

func NewLogger(mux *http.ServeMux) *Logger {
	return &Logger{
		server:  mux,
		slogger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func (l *Logger) Start() app.MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			cw := negroni.NewResponseWriter(w)
			r = r.WithContext(context.WithValue(r.Context(), "smllnk_err", nil))
			next(cw, r)
			start := time.Now()
			err_val := r.Context().Value("smllnk_err")
			if err_val != nil {
				err := err_val.(error)
				l.slogger.Error(err.Error(),
					"route", r.URL.EscapedPath(),
					"method", r.Method,
					"time_taken_s", time.Since(start).Seconds(),
					"time_taken_ms", time.Since(start).Milliseconds(),
					"time_taken_nano", time.Since(start).Nanoseconds(),
					"status", cw.Status(),
					"user_agent", r.UserAgent(),
				)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			l.slogger.Info("success",
				"route", r.URL.EscapedPath(),
				"method", r.Method,
				"time_taken_s", time.Since(start).Seconds(),
				"time_taken_ms", time.Since(start).Milliseconds(),
				"time_taken_nano", time.Since(start).Nanoseconds(),
				"status", cw.Status(),
				"user_agent", r.UserAgent())
		}
	}
}
