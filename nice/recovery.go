package nice

import (
	"io"
	"net/http/httputil"
	
	"github.com/gin-gonic/gin"
)

func Recovery(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return RecoveryWithWriter(f, gin.DefaultErrorWriter)
}

func RecoveryWithWriter(f func(c *gin.Context, err interface{}), out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httprequest, _ := httputil.DumpRequest(c.Request, false)
				reset := string([]byte{27, 91, 48, 109})

				f(c, err)
			}
		}()
		c.Next() // execute all the handlers
	}
}
