package network

import (
	"../application"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func Server(server application.Process, n int, maxDelay int, minDelay int,
	messages chan application.Message) {
	//input: the network# and the # of connections it will receive
	//listen to the client and decode the application, then send via channel
	//simulate the delay here.
	var counter = 0
	ln, err := net.Listen("tcp", server.Ip+":"+server.Port) //creates server
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		if counter == n {
			err = c.Close()
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		decoder := gob.NewDecoder(c)

		mes := new(application.Message)
		_ = decoder.Decode(mes)
		r := rand.Float64() * float64((maxDelay-minDelay)+minDelay)
		time.Sleep(time.Duration(r) * time.Millisecond)
		t := time.Now().Format("Jan _2 15:04:05.000")
		mes.T = t
		messages <- *mes
		counter++
	}
}
