package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type Logger struct {
	server  *echo.Echo
	slogger *slog.Logger
}

func NewLogger(app *echo.Echo) *Logger {
	return &Logger{
		server:  app,
		slogger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func (l *Logger) Start() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			if err := next(c); err != nil {
				l.slogger.Error(err.Error(),
					"method", c.Request().Method,
					"time_taken_s", time.Since(start).Seconds(),
					"time_taken_ms", time.Since(start).Milliseconds(),
					"time_taken_nano", time.Since(start).Nanoseconds(),
					"status", c.Response().Status,
					"user_agent", c.Request().UserAgent(),
				)
				c.Error(err)
			}
			l.slogger.Info("success",
				"method", c.Request().Method,
				"time_taken_s", time.Since(start).Seconds(),
				"time_taken_ms", time.Since(start).Milliseconds(),
				"time_taken_nano", time.Since(start).Nanoseconds(),
				"status", c.Response().Status,
				"user_agent", c.Request().UserAgent())
			return nil
		}
	}
}
