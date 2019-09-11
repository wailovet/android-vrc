package helper

import (
	"encoding/binary"
	"fmt"
	"github.com/wailovet/easycmd"
	"log"
	"strconv"
	"strings"
	"time"
)

func Keyevent(code string, isLongpress bool) {
	if !isLongpress {
		easycmd.EasyCmdNotPty("bash", "-c", "input keyevent "+code).Start(func(data []byte) {
			log.Println(data)
		})
	} else {
		easycmd.EasyCmdNotPty("bash", "-c", "input keyevent --longpress "+code).Start(func(data []byte) {
			log.Println(data)
		})
	}
}

func Swipe(x int, y int, x2 int, y2 int) {
	easycmd.EasyCmdNotPty("bash", "-c", fmt.Sprintf("input keyevent swipe %d %d %d %d", x, y, x2, y2)).Start(func(data []byte) {
		//log.Println(data)
	})
}

type EventData struct {
	Time   int64
	Device string
	Data   InputEvent
}

func SendEvent(text string) {
	var evendata []EventData
	lines := strings.Split(text, "\n")
	for e := range lines {
		line := strings.Split(lines[e], " ")
		if len(line) != 6 {
			log.Printf("format error:%v", line)
			continue
		}

		timee := line[0]
		device := line[2]

		typet, _ := strconv.Atoi(line[3])
		code, _ := strconv.Atoi(line[4])
		value, _ := strconv.Atoi(line[5])

		floatmp, _ := strconv.ParseFloat(timee, 64)
		floatmp = floatmp * float64(time.Millisecond)
		tt := int64(time.Second/time.Millisecond) * int64(floatmp)
		if floatmp == 0 {
			tt = 0
		}
		_, _ = getDevice(device)
		evendata = append(evendata, EventData{
			Time:   tt,
			Device: device,
			Data: InputEvent{
				Type:  uint16(typet),
				Code:  uint16(code),
				Value: int32(value),
			},
		})
	}

	for e := range evendata {
		if evendata[e].Time > 0 {
			time.Sleep(time.Duration(evendata[e].Time))
		}
		tf, _ := getDevice(evendata[e].Device)

		err := binary.Write(tf, binary.LittleEndian, &evendata[e].Data)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
