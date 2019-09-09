package helper

import (
	"github.com/wailovet/easycmd"
	"log"
)

func Keyevent(code string, isLongpress bool) {
	if !isLongpress {
		easycmd.EasyCmd("bash", "-c", "input keyevent "+code).Start(func(data []byte) {
			log.Println(data)
		})
	} else {
		easycmd.EasyCmd("bash", "-c", "input keyevent --longpress "+code).Start(func(data []byte) {
			log.Println(data)
		})
	}
}
