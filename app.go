package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	jackett  *Jackett
	database *Database
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Load Jackett settings from localStorage via database settings
	if a.jackett != nil && a.database != nil {
		settings, err := a.database.GetSettings()
		if err == nil {
			// Apply Jackett settings if available
			if settings.JackettURL != "" {
				a.jackett.SetConfig(settings.JackettURL, settings.JackettAPIKey)
			}
		}
	}
}

// shutdown is called when the app is shutting down
func (a *App) shutdown(ctx context.Context) {
	// Cleanup code here
}

// WindowMinimize minimizes the window
func (a *App) WindowMinimize() {
	runtime.WindowMinimise(a.ctx)
}

// WindowMaximize toggles window maximize
func (a *App) WindowMaximize() {
	runtime.WindowToggleMaximise(a.ctx)
}

// WindowClose closes the window
func (a *App) WindowClose() {
	runtime.Quit(a.ctx)
}
