package terminal

import (
	// "bytes"
	// "fmt"
	// "io"
	// "net/http"
	// "log"
	//"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func StartTerminalSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		newSession := New("ls -lah")


		stdout, err := newSession.Command.StdoutPipe()
		buffer := new(bytes.Buffer)

	}
}

func SendTerminalCharacters() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
