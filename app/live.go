package app

import (
	"encoding/base64"
	"github.com/wailovet/android-vrc/helper"
	"github.com/wailovet/easycmd"
	"github.com/wailovet/osmanthuswine/src/core"
	"gopkg.in/olahol/melody.v1"
	"log"
)

type Live struct {
	core.WebSocket
}

func (that *Live) HandleConnect(s *melody.Session) {
	//panic("implement me")
	s.Set("c", helper.TakeScreenrecord(func(bytes []byte) {
		log.Println("len:", len(bytes))
		data := base64.StdEncoding.EncodeToString(bytes)
		s.Write([]byte(data))
	}))
}

func (that *Live) HandlePong(*melody.Session) {
	//panic("implement me")
}

func (that *Live) HandleMessage(*melody.Session, []byte) {
	//panic("implement me")
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
