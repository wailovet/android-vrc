package app

import (
	"encoding/base64"
	"encoding/json"
	"github.com/wailovet/android-vrc/helper"
	"github.com/wailovet/easycmd"
	"github.com/wailovet/osmanthuswine/src/core"
	"gopkg.in/olahol/melody.v1"
	"log"
	"strconv"
	"strings"
)

type Live struct {
	core.WebSocket
	w int
	h int
}

func (that *Live) HandleConnect(s *melody.Session) {
	//panic("implement me")
	//that.w, that.h = helper.GetWmSize()
	that.w = 0x3FF
	that.h = 0x3FF
	s.Set("c", helper.TakeScreenrecord(func(bytes []byte) {
		//log.Println("len:", len(bytes))
		data := base64.StdEncoding.EncodeToString(bytes)
		s.Write([]byte(data))
	}))
}

func (that *Live) HandlePong(*melody.Session) {
	//panic("implement me")
}

func (that *Live) HandleMessage(s *melody.Session, data []byte) {
	//panic("implement me")
	type event struct {
		Type string `json:"type"`
		Data string `json:"data"`
	}

	e := event{}
	json.Unmarshal(data, &e)

	log.Println(e.Type, ":", e.Data)
	switch e.Type {
	case "key":
		helper.Keyevent(e.Data, false)
		break
	case "keylong":
		helper.Keyevent(e.Data, true)
		break
	case "touchstart":
		touchData := strings.Split(e.Data, ",")
		x1, _ := strconv.ParseFloat(touchData[0], 64)
		y1, _ := strconv.ParseFloat(touchData[1], 64)
		helper.TouchDown(x1, y1, 1)
		break
	case "touchmove":
		touchData := strings.Split(e.Data, ",")
		x1, _ := strconv.ParseFloat(touchData[0], 64)
		y1, _ := strconv.ParseFloat(touchData[1], 64)
		helper.TouchMove(x1, y1, 1)
		break
	case "touchend":
		helper.TouchUp(1)
		break
	}
}

func (that *Live) HandleMessageBinary(*melody.Session, []byte) {
	//panic("implement me")
}

func (that *Live) HandleSentMessage(*melody.Session, []byte) {
	//panic("implement me")
}

func (that *Live) HandleSentMessageBinary(*melody.Session, []byte) {
	//panic("implement me")
}

func (that *Live) HandleDisconnect(s *melody.Session) {
	//panic("implement me")
	log.Println("HandleDisconnect")
	c, exists := s.Get("c")
	if exists {
		c.(*easycmd.Pty).Close()
	}
}

func (that *Live) HandleError(*melody.Session, error) {
	//panic("implement me")
}
