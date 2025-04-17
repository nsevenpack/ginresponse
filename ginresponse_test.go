package ginresponse

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSuccessResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/test", nil)

	SetFormatter(&JsonFormatter{})

	Success(c, "Everything is OK", map[string]string{"foo": "bar"})

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), `"message":"Everything is OK"`)
	assert.Contains(t, w.Body.String(), `"foo":"bar"`)
}

func TestCreatedResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/test", nil)

	SetFormatter(&JsonFormatter{})

	Created(c, "Resource created", map[string]string{"id": "123"})

	assert.Equal(t, http.StatusCreated, w.Code)

	assert.Contains(t, w.Body.String(), `"message":"Resource created"`)
	assert.Contains(t, w.Body.String(), `"id":"123"`)
}

func TestErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/test", nil)

	SetFormatter(&JsonFormatter{})

	Error(c, http.StatusBadRequest, "Bad Request", nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	assert.Contains(t, w.Body.String(), `"message":"Bad Request"`)
}

func TestSetFormatter(t *testing.T) {
	SetFormatter(&JsonFormatter{})
	assert.NotNil(t, formatter)

	SetFormatter(nil)
	assert.Nil(t, formatter)

	// 3. On teste que ça panique bien si on appelle un helper sans formatter
	assert.Panics(t, func() {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/panic", nil)

		Success(c, "no formatter", nil)
	})
}