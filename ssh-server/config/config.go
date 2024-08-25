package config

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// expect ssh_host_key, ssh_client_key and chess.db to be mounted in the container here
const (
	DbPath = "file:/data/chess.db?cache=shared&mode=rwc"

	SshServerAddress     = "0.0.0.0:22"
	ShhServerHostKeyPath = "/certs/ssh_host_key"

	HttpServerAddress = "0.0.0.0:80"
	SshClientKeyPath  = "/certs/ssh_client_key"
)

const (
	MinWidth  = 40
	MinHeight = 24
)

var (
	TeaOptions = []tea.ProgramOption{tea.WithAltScreen(), tea.WithOutput(os.Stderr)}
)
