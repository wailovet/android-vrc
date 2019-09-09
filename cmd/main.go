package main

import (
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/wailovet/android-vrc/app"
	"github.com/wailovet/osmanthuswine"
	"github.com/wailovet/osmanthuswine/src/core"
)

func main() {

	core.GetInstanceConfig().Port = "25009"
	core.GetInstanceConfig().Host = "0.0.0.0"
	core.GetInstanceConfig().StaticRouter = "/*"
	core.GetInstanceConfig().StaticFileSystem = &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "../static"}

	core.GetInstanceRouterManage().Registered(&app.Live{})
	osmanthuswine.Run()

}
