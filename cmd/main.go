package main

import (
	"github.com/wailovet/android-vrc/app"
	"github.com/wailovet/android-vrc/helper"
	"github.com/wailovet/osmanthuswine"
	"github.com/wailovet/osmanthuswine/src/core"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

func main() {

	//go-bindata-assetfs static/..
	core.GetInstanceConfig().Port = "25009"
	core.GetInstanceConfig().Host = "0.0.0.0"
	core.GetInstanceConfig().StaticFileSystem = assetFS()
	osmanthuswine.GetChiRouter().HandleFunc("/screenrecord.mp4", func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Add("Content-Type", "audio/mp4")
		helper.TakeScreenrecord(func(bytes []byte) {
			writer.Write(bytes)
		})
	})
	osmanthuswine.GetChiRouter().HandleFunc("/snapshot.jpg", func(writer http.ResponseWriter, request *http.Request) {

		img, err := helper.TakeSnapshot()
		if err != nil {
			println(err.Error())
			return
		}
		writer.Header().Add("Content-Type", "image/jpeg")
		q := request.FormValue("q")
		log.Println("q:", q)
		iq, err := strconv.Atoi(q)
		if err != nil {
			iq = 50
		}
		jpeg.Encode(writer, img, &jpeg.Options{Quality: iq})
	})
	core.GetInstanceRouterManage().Registered(&app.Live{})
	osmanthuswine.Run()

}
