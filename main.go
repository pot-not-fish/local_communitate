package main

import (
	"embed"
	"my_local_communitate/pkg/cache/group"
	"my_local_communitate/pkg/cache/lru"

	// "my_local_communitate/pkg/cache/group"
	// "my_local_communitate/pkg/cache/lru"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "my_local_communitate",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

	lruCache := lru.NewCache()
	group.NewGroup("symmetric_key", lruCache)
	group.NewGroup("asymmetric_key", lruCache)

	// web server
	r := gin.Default()

	r.POST("/upload")
	r.POST("/keygen")

	r.Run("0.0.0.0:5000")
}
