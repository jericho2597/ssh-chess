package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sshServer := setupSSHServer()
	setupHTTPServer()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go startSshServer(sshServer)
	go startHTTPServer()

	<-done
	shutdownSshServer(sshServer)
}
