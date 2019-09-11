package main

import (
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/wailovet/android-vrc/app"
	"github.com/wailovet/android-vrc/helper"
	"github.com/wailovet/osmanthuswine"
	"github.com/wailovet/osmanthuswine/src/core"
	"log"
)

func main() {
	core.SetConfig(&core.Config{
		Host:         "0.0.0.0",
		Port:         "25009",
		ApiRouter:    "/Api/*",
		StaticRouter: "/*",
	})

	//自动检测android可用的临时目录
	tmpDir := helper.AutoSetTmpDir()
	if tmpDir == "" {
		log.Println("上传文件不可用")
	} else {
		log.Println("临时文件目录:", tmpDir)
	}

	helper.GetWmSize()
	core.GetInstanceConfig().StaticFileSystem = &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "../static"}
	core.GetInstanceRouterManage().Registered(&app.Live{})
	osmanthuswine.Run()

}
