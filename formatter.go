package ginresponse

import "github.com/gin-gonic/gin"

type Formatter interface {
	Format(c *gin.Context, status int, message string, data interface{}, err interface{}) map[string]interface{}
	Render(c *gin.Context, status int, payload map[string]interface{})
}
 
var formatter Formatter

func SetFormatter(f Formatter) {
	formatter = f
}

func ensureFormatterSet() {
	if formatter == nil {
		panic("ginresponse: no formatter set. You must call ginresponse.SetFormatter(...) before using response helpers.")
	}
}