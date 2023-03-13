package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).Info("incoming request")
		return next(c)
	}
}
func makeLogEntry(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return logrus.WithFields(logrus.Fields{
		"at":         time.Now().Format("2006-01-02 15:04:05"),
		"method":     c.Request().Method,
		"uri":        c.Request().URL.String(),
		"ip":         c.Request().RemoteAddr,
		"user_agent": c.Request().UserAgent(),
	})
}
