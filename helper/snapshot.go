package helper

import (
	"bytes"
	"encoding/binary"
	"github.com/wailovet/easycmd"
	"image"
	"os/exec"
)

var SCRBUFLEN int

func cmdSnapshot() *exec.Cmd {

	return exec.Command("screencap")
}

func TakeScreenrecord(f func([]byte)) *easycmd.Pty {
	c := easycmd.EasyCmd("screenrecord", "--output-format=h264", "--bit-rate=16m", "--size=800x600", "-")
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
func TakeSnapshot() (img image.Image, err error) {
	var scrbf *bytes.Buffer
	if SCRBUFLEN == 0 {
		scrbf = bytes.NewBuffer(nil)
	} else {
		scrbf = bytes.NewBuffer(make([]byte, 0, SCRBUFLEN))
	}
	var cmd *exec.Cmd

	cmd = cmdSnapshot()
	cmd.Stdout = scrbf
	if err = cmd.Run(); err != nil {
		return
	}
	var width, height, format int32
	binary.Read(scrbf, binary.LittleEndian, &width)
	binary.Read(scrbf, binary.LittleEndian, &height)
	SCRBUFLEN = int(width * height * 4)
	err = binary.Read(scrbf, binary.LittleEndian, &format)
	if err != nil {
		return
	}
	w, h := int(width), int(height)
	var buf []byte

	buf = scrbf.Bytes()

	img = &image.RGBA{
		Pix:    buf,
		Stride: 4 * w,
		Rect:   image.Rect(0, 0, w, h),
	}
	m := img
	//m := resize.Resize(300, 0, img, resize.Bilinear)

	return m, err
}
