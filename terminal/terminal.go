package terminal

import (
	"math/rand/v2"
	"os/exec"
	"strings"
	// "time"
)

type TerminalSession struct {
	SessionID int
	InitCommand string
	Command *exec.Cmd
	InputEventChannel chan string
}

func New(command string) TerminalSession {
	sessionId := genSessionID()
	initCommand := command

	command_args := strings.Split(initCommand, " ");
	cmd := exec.Command(command_args[0], command[1:]); //Remember you can get the stdin and stdout from this guy

	var output TerminalSession;
	output.SessionID = sessionId
	output.Command = cmd
	output.InitCommand = command
	output.InputEventChannel = make(chan string)
	return output
}

func genSessionID() int {

	return rand.IntN(25565)
}

func (term *TerminalSession) Destroy() {

}
