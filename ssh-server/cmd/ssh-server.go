package main

import (
	"context"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"

	"ssh-server/config"
	"ssh-server/internal/app"
)

func setupSSHServer() *ssh.Server {

	log.Printf("Attempting to load host key from: %s", config.ShhServerHostKeyPath)
	hostKey, err := os.ReadFile(config.ShhServerHostKeyPath)
	if err != nil {
		log.Fatalf("Failed to load host key: %v", err)
	}
	log.Printf("Host key loaded successfully, length: %d bytes", len(hostKey))

	server, err := wish.NewServer(
		wish.WithAddress(config.SshServerAddress),
		wish.WithHostKeyPath(config.ShhServerHostKeyPath),
		wish.WithMiddleware(
			bm.Middleware(chessMiddleware),
			lm.Middleware(),
		),
	)

	if err != nil {
		log.Fatalf("Could not create SSH server: %s", err)
	}
	return server
}

func startSshServer(server *ssh.Server) {
	log.Printf("Starting SSH server on %s", config.SshServerAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting SSH server: %s", err)
	}
}

func shutdownSshServer(server *ssh.Server) {
	log.Println("stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
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

	tui := app.NewTUI(pty.Window.Width, pty.Window.Height)

	go func() {
		<-s.Context().Done()
	}()

	return tui, config.TeaOptions
}
