package view

import (
	"fmt"
	"ssh-server/pkg/app"
	"time"
)

func RenderLoadingScreen(tui *app.TUI) string {

	app.SchedulerService.Schedule(tui.SessionId, "", 1*time.Second)

	t := time.Now()
	return fmt.Sprintf("%s tui.SessionId", t.String())
}
