package ginresponse

import (
	"time"

	"github.com/gin-gonic/gin"
)

type JsonFormatter struct{}

type EmptyObject map[string]interface{}

type JsonFormatterSwag struct {
	Message string       `json:"message"`
	Data    EmptyObject  `json:"data"`
	Error   EmptyObject  `json:"error"`
	Meta    JsonSwagMeta `json:"meta"`
}

type JsonSwagMeta struct {
	Status    int    `json:"status"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	Timestamp string `json:"timestamp"`
}

func (f *JsonFormatter) Format(c *gin.Context, status int, message string, data interface{}, err interface{}) map[string]interface{} {
	resp := map[string]interface{}{
		"message": message,
		"data":    nil,
		"error":   nil,
		"meta": map[string]interface{}{
			"status":    status,
			"path":      c.Request.URL.Path,
			"method":    c.Request.Method,
			"timestamp": time.Now().Format(time.RFC3339),
		},
	}

	if data != nil {
		resp["data"] = data
	}
	if err != nil {
		resp["error"] = err
	}

	return resp
}

func (f *JsonFormatter) Render(c *gin.Context, status int, payload map[string]interface{}) {
	c.JSON(status, payload)
}
