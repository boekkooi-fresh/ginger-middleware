package nice

import (
	"io"
	"errors"
	"net/http/httputil"

	log "github.com/Sirupsen/logrus"
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
				goErr := errors.Wrap(err, 3)
				reset := string([]byte{27, 91, 48, 109})
				log.Debugf("[Nice Recovery] panic recovered:\n\n%s%s\n\n%s%s", httprequest, goErr.Error(), goErr.Stack(), reset)

				f(c, err)
			}
		}()
		c.Next() // execute all the handlers
	}
}
