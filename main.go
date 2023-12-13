package main

import (
	"embed"
	"fmt"

	"fileuber/server"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

var scfg = server.ShareCfg


func main() {
	// Create an instance of the app structure
	app := NewApp()
	
	localserver := server.NewLocalServer()
	
    
	
	

	// Create application with options
	err := wails.Run(

		&options.App{
		EnableDefaultContextMenu: false,

		Width: 760 ,
		Height: 480,
		DisableResize: true,	
		Title:  "fileuber",
		LogLevel: logger.ERROR,
		
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			IsZoomControlEnabled: false,
			ZoomFactor: 1.0,

		},
		SingleInstanceLock: &options.SingleInstanceLock{
			UniqueId: "fileuberappinstance",
			OnSecondInstanceLaunch: func(secondInstanceData options.SecondInstanceData) {
				fmt.Print("Another instance launched")
			},
		},
		
		// rgba(197, 239, 203, 1)
		BackgroundColour: &options.RGBA{R: 197, G: 239, B: 203, A: 1},
		OnStartup:        app.startup,

		Bind: []interface{}{
			app,
			scfg,
			localserver,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
	
}
