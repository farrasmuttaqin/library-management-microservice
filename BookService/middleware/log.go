package middleware

import (
	"book_service/models"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"runtime"
	"time"
)

// AccessLogger logger for splunk
func AccessLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Capture request payload
		var payloadMap interface{}
		if err := json.Unmarshal(c.Body(), &payloadMap); err != nil {
			payloadMap = map[string]interface{}{}
		}

		// Process the request
		err := c.Next()
		end := time.Now()

		latency := end.Sub(start)

		microServiceName := "book_service"

		// Convert response body to map[string]interface{}
		//var responseMap map[string]interface{}
		//if err2 := json.Unmarshal(rec.body.Bytes(), &responseMap); err2 != nil {
		//	responseMap = map[string]interface{}{
		//		"error": "failed to unmarshal response",
		//	}
		//}

		fullPath := c.Request().URI().String()
		routePath := c.Route().Path
		if routePath == "" {
			routePath = fullPath // Fallback to fullPath if routePath is empty
		}

		// Prepare log data
		logData := map[string]interface{}{
			"MICROSERVICE_NAME":    microServiceName,
			"CLIENT_IP":            c.IP(),
			"DATE_TIME":            end.Format("02/Jan/2006:15:04:05"),
			"METHOD":               c.Method(),
			"PATH":                 routePath,
			"RESPONSE_CODE":        c.Response().StatusCode(),
			"LATENCY_MILLISECONDS": latency.Milliseconds(),
			"LATENCY_SECONDS":      latency.Seconds(),
			"USER_AGENT":           c.Request().Header.UserAgent(),
			"HEADER":               c.Request().Header,
			"REQUEST_BODY":         payloadMap,
		}
		// Print log to console
		logJSON, err2 := json.Marshal(logData)
		if err2 != nil {
			log.Println(err2)
		}
		log.Println(string(logJSON))

		return err
	}
}

// ErrorHandler logger for error
func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a recovery middleware
		defer func() {
			if r := recover(); r != nil {
				// Get the panic details
				err := fmt.Errorf("recovered from panic: %v", r)

				// Get the path
				path := c.Path()
				if path == "" {
					path = "unknown"
				}

				// Get the query params as a string
				queryParams := c.Context().QueryArgs().String() // Retrieve query parameters as a string
				if queryParams == "" {
					queryParams = "unknown"
				}

				// Get the request body
				reqBody := fmt.Sprintf("%v", string(c.Body())) // Use string conversion for readable output

				// Get the caller information
				pc, file, line, ok := runtime.Caller(1)
				if !ok {
					// If caller information is not available, set default values
					file = "???"
					line = 0
				}

				// Get the function name
				functionName := runtime.FuncForPC(pc).Name()

				// Get the stack trace
				stackTraceBuf := make([]byte, 4096)
				stackTraceLen := runtime.Stack(stackTraceBuf, false)
				stackTrace := string(stackTraceBuf[:stackTraceLen])

				mapErrorRequest := models.MapErrorRequest{
					QueryParams: queryParams,
					RequestBody: reqBody,
				}

				mapErrorResponse := models.MapErrorResponse{
					Path:         path,
					Err:          err.Error(),
					FunctionName: functionName,
					File:         file,
					Line:         line,
					StackTrace:   stackTrace,
				}

				// Log the error with all details
				logData := map[string]interface{}{
					"LOG":      "LOG_ERROR_AUTHOR_MICROSERVICES",
					"REQUEST":  mapErrorRequest,
					"RESPONSE": mapErrorResponse,
				}
				logJSON, _ := json.Marshal(logData)
				log.Println(string(logJSON))

				// Set an internal server error response
				c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": 0, "error": "Internal server error"})
			}
		}()

		// Continue to the next handler
		return c.Next()
	}
}
