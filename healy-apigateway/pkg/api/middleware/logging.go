package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"healy-apigateway/pkg/logging"
)

func LoggingMiddleware(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		logger := logging.Logger().WithFields(logrus.Fields{
			"method": c.Method(),
			"path":   c.Path(),
			"ip":     c.IP(),
		})

		logEntry := logger.WithFields(logrus.Fields{
			"request": c.OriginalURL(),
		})

		logEntry.Info("Incoming request")

		err := next(c)

		duration := time.Since(start)
		var statusCode int
		if c.Response() != nil {
			statusCode = c.Response().StatusCode()
		}
		if err != nil || (statusCode >= 400 && statusCode < 500) {
			errStr := ""
			if err != nil {
				errStr = err.Error()
			}
			logEntry.WithFields(logrus.Fields{
				"status":   statusCode,
				"duration": duration.Milliseconds(),
				"error":    errStr,
			}).Error("Request failed")
		} else {
			logEntry.WithFields(logrus.Fields{
				"status":   statusCode,
				"duration": duration.Milliseconds(),
			}).Info("Completed request")
		}

		return err
	}
}
