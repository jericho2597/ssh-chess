package config

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// expect ssh_host_key and chess.db to be mounted in the container here
const (
	HostKeyPath = "/app/ssh_host_key"
	DbPath      = "file:/data/chess.db?cache=shared&mode=rwc"
)

const (
	MinWidth  = 40
	MinHeight = 24
)

var (
	TeaOptions = []tea.ProgramOption{tea.WithAltScreen(), tea.WithOutput(os.Stderr)}
)
