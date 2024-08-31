package main

import (
	"context"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/wish"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"

	"ssh-server/config"
	"ssh-server/pkg/app"
	"ssh-server/pkg/app/update"
	"ssh-server/pkg/app/view"

	"github.com/muesli/termenv"
)

func setupSSHServer() *ssh.Server {

	log.Printf("Attempting to load host key from: %s", config.ShhServerHostKeyPath)
	hostKey, err := os.ReadFile(config.ShhServerHostKeyPath)
	if err != nil {
		log.Fatalf("Failed to load host key: %v", err)
	}
	log.Printf("Host key loaded successfully, length: %d bytes", len(hostKey))

	bubbleteaMiddleware := MiddlewareWithColorProfile(ChessMiddleware, termenv.ANSI)

	server, err := wish.NewServer(
		wish.WithAddress(config.SshServerAddress),
		wish.WithHostKeyPath(config.ShhServerHostKeyPath),
		wish.WithMiddleware(
			bubbleteaMiddleware,
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("could not shutdown server gracefully: %s", err)
	}
}

func ChessMiddleware(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		log.Printf("no active terminal, skipping")
		return nil, nil
	}

	viewer := view.View{}
	updater := update.Update{}

	sessionId := s.Context().SessionID()
	tui := app.NewTUI(sessionId, viewer, updater, pty.Window.Width, pty.Window.Height)

	go func() {
		<-s.Context().Done()
		app.SessionManagerService.RemoveSession(s.Context().SessionID())
	}()

	return tui, config.TeaOptions
}

type Handler func(ssh.Session) (tea.Model, []tea.ProgramOption)

func MiddlewareWithColorProfile(bth Handler, cp termenv.Profile) wish.Middleware {

	return func(sh ssh.Handler) ssh.Handler {
		lipgloss.SetColorProfile(cp)
		return func(s ssh.Session) {
			errc := make(chan error, 1)
			m, opts := bth(s)
			if m != nil {
				opts = append(opts, tea.WithInput(s), tea.WithOutput(s))
				p := tea.NewProgram(m, opts...)
				_, windowChanges, _ := s.Pty()
				go func() {
					for {
						select {
						case <-s.Context().Done():
							if p != nil {
								p.Quit()
							}
						case w := <-windowChanges:
							if p != nil {
								p.Send(tea.WindowSizeMsg{Width: w.Width, Height: w.Height})
							}
						case err := <-errc:
							if err != nil {
								log.Print(err)
							}
						}
					}
				}()

				app.SessionManagerService.AddSession(s, *p)

				errc <- p.Start()
			}
			sh(s)
		}
	}
}
