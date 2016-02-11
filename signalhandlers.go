package signalhandlers

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Interrupt blocking listener that exits with success (0) when SIGINT is found.
func Interrupt() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	fmt.Println("")
	log.Println("CTRL-C (SIGINT); exiting")
	os.Exit(0)
}

// Quit blocking listener that exits with failure (1) when SIGQUIT is found.
func Quit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGQUIT)
	<-ch
	fmt.Println("")
	log.Println("CTRL-\\ (SIGQUIT); exiting")
	os.Exit(1)
}
