package ginresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusOK, message, data, nil)
	formatter.Render(c, http.StatusOK, payload)
}
func Created(c *gin.Context, message string, data interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusCreated, message, data, nil)
	formatter.Render(c, http.StatusCreated, payload)
}
func NoContent(c *gin.Context, message string) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusNoContent, message, nil, nil)
	formatter.Render(c, http.StatusNoContent, payload)
}
func BadRequest(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusBadRequest, message, nil, err)
	formatter.Render(c, http.StatusBadRequest, payload)
}
func Unauthorized(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusUnauthorized, message, nil, err)
	formatter.Render(c, http.StatusUnauthorized, payload)
}
func Forbidden(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusForbidden, message, nil, err)
	formatter.Render(c, http.StatusForbidden, payload)
}
func NotFound(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusNotFound, message, nil, err)
	formatter.Render(c, http.StatusNotFound, payload)
}
func Conflict(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusConflict, message, nil, err)
	formatter.Render(c, http.StatusConflict, payload)
}
func InternalServerError(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusInternalServerError, message, nil, err)
	formatter.Render(c, http.StatusInternalServerError, payload)
}
func ServiceUnavailable(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusServiceUnavailable, message, nil, err)
	formatter.Render(c, http.StatusServiceUnavailable, payload)
}
func UnprocessableEntity(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusUnprocessableEntity, message, nil, err)
	formatter.Render(c, http.StatusUnprocessableEntity, payload)
}
func TooManyRequests(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusTooManyRequests, message, nil, err)
	formatter.Render(c, http.StatusTooManyRequests, payload)
}
func NotAcceptable(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusNotAcceptable, message, nil, err)
	formatter.Render(c, http.StatusNotAcceptable, payload)
}
func MethodNotAllowed(c *gin.Context, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, http.StatusMethodNotAllowed, message, nil, err)
	formatter.Render(c, http.StatusMethodNotAllowed, payload)
}
func Error(c *gin.Context, status int, message string, err interface{}) {
	ensureFormatterSet()
	payload := formatter.Format(c, status, message, nil, err)
	formatter.Render(c, status, payload)
}