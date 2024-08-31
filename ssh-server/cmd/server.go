package main

import (
	"os"
	"os/signal"
	"ssh-server/pkg/app"
	"syscall"
)

func main() {

	app.NewSessionManager()
	app.NewScheduler(app.SessionManagerService)

	sshServer := setupSSHServer()
	setupHTTPServer()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go startSshServer(sshServer)
	go startHTTPServer()

	<-done
	app.SchedulerService.Shutdown()
	shutdownSshServer(sshServer)
}
