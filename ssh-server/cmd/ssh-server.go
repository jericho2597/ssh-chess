package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"

	"ssh-server/config"
	"ssh-server/pkg/app"
	"ssh-server/pkg/app/update"
	"ssh-server/pkg/app/view"
)

func main() {

	log.Printf("Attempting to load host key from: %s", config.HostKeyPath)
	hostKey, err := ioutil.ReadFile(config.HostKeyPath)
	if err != nil {
		log.Fatalf("Failed to load host key: %v", err)
	}
	log.Printf("Host key loaded successfully, length: %d bytes", len(hostKey))

	server, err := wish.NewServer(
		wish.WithAddress("0.0.0.0:8080"),
		wish.WithIdleTimeout(5*time.Minute),
		wish.WithHostKeyPath(config.HostKeyPath),
		wish.WithMiddleware(
			bm.Middleware(chessMiddleware),
			lm.Middleware(),
		),
	)
	if err != nil {
		log.Fatalf("could not create server: %s", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("starting server: %s", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("error starting SSH server: %s", err)
		}
	}()

	<-done
	log.Println("stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("could not shutdown server gracefully: %s", err)
	}
}

func chessMiddleware(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		log.Printf("no active terminal, skipping")
		return nil, nil
	}

	viewer := view.View{}
	updater := update.Update{}

	tui := app.NewTUI(viewer, updater, pty.Window.Width, pty.Window.Height)

	return tui, config.TeaOptions
}
