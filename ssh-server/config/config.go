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
	HeaderHeight = 4
	FooterHeight = 4

	ContentMinWidth  = 64 // 8 (thinnest supported square in cells) times 8 (squares along width of board)
	ContentMinHeight = 32 // 4 (shortest supportedsquare in cells) times 8 (squares on along side of board)

	ContentMaxWidth  = 96 // 12 (widest supported square in cells) times 8 (squares along width of board)
	ContentMaxHeight = 48 // 6 (tallest supportedsquare in cells) times 8 (squares on along side of board)

	MinWidth  = ContentMinWidth
	MinHeight = HeaderHeight + ContentMinHeight + FooterHeight
)

var (
	TeaOptions = []tea.ProgramOption{tea.WithAltScreen(), tea.WithOutput(os.Stderr), tea.WithMouseAllMotion()}
)
