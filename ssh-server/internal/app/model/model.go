package model

type Screen int

const (
	GameScreen Screen = iota
	AboutScreen
	LoadingScreen
)

type TuiState struct {
	Height         int
	Width          int
	SufficientSize bool
	ActiveScreen   Screen
}

func NextScreen(currentScreen Screen) Screen {
	s := currentScreen + 1
	if s == LoadingScreen {
		s = GameScreen
	}
	return s
}
