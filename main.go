package main

import (
	"cypher/internal/services"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	credentialSrv := services.NewCredentialService()
	//tools.Notify("Cypher", "启动成功")

	// Create application with options
	err := wails.Run(&options.App{
		Title: "\U0001FABA Cypher",
		//Frameless: true,
		Width:  900,
		Height: 458,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: options.NewRGB(255, 255, 255),
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			credentialSrv,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
