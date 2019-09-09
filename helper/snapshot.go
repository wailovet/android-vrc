package helper

import (
	"fmt"
	"github.com/wailovet/easycmd"
	"log"
	"strconv"
	"strings"
	"time"
)

func GetWmSize() (int, int) {
	c := easycmd.EasyCmd("bash", "-c", "dumpsys window displays |head -n 3 | grep x")
	w := 0
	h := 0

	c.SetEventEnd(func() {
	})

	c.Start(func(data []byte) {
		s := strings.TrimSpace(string(data))
		log.Println("getWmSize:", s)
		s2 := strings.Split(s, " ")
		s2 = strings.Split(s2[0], "=")
		if len(s2) > 1 {
			s3 := strings.Split(s2[1], "x")
			w, _ = strconv.Atoi(strings.TrimSpace(s3[0]))
			h, _ = strconv.Atoi(strings.TrimSpace(s3[1]))
			log.Println("getWmSize:", w, "  |   ", h)
			c.Close()
		}

	})

	for i := 0; i < 5 && w == 0 && h == 0; i++ {
		time.Sleep(time.Second)
	}
	return w, h
}
func TakeScreenrecord(f func([]byte)) *easycmd.Pty {
	w, h := GetWmSize()
	sf := h / 600
	size := fmt.Sprintf("--size=%dx%d", w/sf, h/sf)
	log.Println("size:", size)
	c := easycmd.EasyCmd("screenrecord", "--output-format=h264", "--bit-rate=500000", size, "-")
	//var ch chan bool
	//c.SetEventEnd(func() {
	//	ch <- true
	//})
	c.Start(func(data []byte) {
		f(data)
	})
	return c
	//<-ch
}
