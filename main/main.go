package main
import (
	"pkg/display"
	"pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("Hello from display")
	msg.Exciting("An exciting message")
	
}

