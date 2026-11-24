package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	if err := wails.Run(&options.App{
		Title:  "PhongDolphin",
		Width:  650,
		Height: 500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
	}); err != nil {
		println("Error:", err.Error())
	}
}
