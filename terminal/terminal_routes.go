package terminal

import (
	// "bytes"
	// "fmt"
	// "io"
	"io"
	"net/http"
	// "log"
	//"os"
	// "sync"
	"bufio"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartTerminalSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		connection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Upgrade:", err)
			return
		}
		defer connection.Close()

		newSession := New("ls -lah /etc")

		termoutput, err := newSession.Command.StdoutPipe()
		if err != nil {
			log.Println("Error Getting Stdout Pipe: ", err)
			return
		}
		termerr, err := newSession.Command.StderrPipe()
		if err != nil {
			log.Println("Error Getting Stderr Pipe: ", err)
			return
		}


		defer termoutput.Close()
		defer termerr.Close()


		errStart := newSession.Command.Start()
		log.Println(newSession.Command.String())

		if errStart != nil {
			log.Println("Error Running command: ", err);
		}
		go readPipe(termoutput, connection)

		errEnd := newSession.Command.Wait()

		if errEnd != nil {
			log.Println("Error Finishing Command: ", errEnd);
		}


	}
}

func readPipe(inputReader io.ReadCloser, socket *websocket.Conn) {
	pipeScanner := bufio.NewScanner(inputReader)
	for pipeScanner.Scan() {
		//message := stdoutScanner.Text()
		//log.Println("Text: ", pipeScanner.Text())
		err := socket.WriteMessage(websocket.TextMessage, pipeScanner.Bytes())
		if err != nil {
			log.Println("Error Writing Message", err)
		}
		if err := pipeScanner.Err(); err != nil {
			if err != io.EOF {
				log.Println("Error Reading from out pipe: ", err)
			}
		}
	}

}
