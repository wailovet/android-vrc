package helper

import (
	"fmt"
	"github.com/wailovet/easycmd"
	"log"
	"strconv"
	"strings"
)

func GetWmSize() (int, int) {
	c := easycmd.EasyCmd("dumpsys", "window", "displays")
	ch := make(chan bool)
	w := 0
	h := 0

	c.SetEventEnd(func() {
		log.Println("GetWmSize,end")
		ch <- true
	})

	c.Start(func(data []byte) {
		ss := strings.Split("\n", string(data))
		for e := range ss {
			s := strings.TrimSpace(ss[e])
			log.Println("getWmSize:", s)
			s2 := strings.Split(s, " ")
			s2 = strings.Split(s2[0], "=")
			if len(s2) > 1 {
				s3 := strings.Split(s2[1], "x")
				if len(s3) > 1 {
					wt, err := strconv.Atoi(strings.TrimSpace(s3[0]))
					ht, err2 := strconv.Atoi(strings.TrimSpace(s3[1]))
					if err != nil || err2 != nil || ht == 0 || wt == 0 {
						return
					}
					w = wt
					h = ht
					c.Close()
				}
			}
		}

	})

	<-ch

	log.Println("getWmSize:", w, "  |   ", h)
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
