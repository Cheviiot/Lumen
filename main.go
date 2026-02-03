package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed Icons/Lumen(1).png
var icon []byte

func main() {
	// Create application with options
	app := NewApp()
	torrServer := NewTorrServer()
	jackett := NewJackett()
	tmdb := NewTMDB()
	database := NewDatabase()

	// Link instances to app for initialization
	app.jackett = jackett
	app.database = database

	err := wails.Run(&options.App{
		Title:             "Lumen",
		Width:             1280,
		Height:            800,
		MinWidth:          1024,
		MinHeight:         600,
		Frameless:         true,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 18, G: 18, B: 18, A: 255},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
			torrServer,
			jackett,
			tmdb,
			database,
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         "Lumen",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
