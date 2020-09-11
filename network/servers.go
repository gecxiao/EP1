package network

import (
	"../application"
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func Server(server application.Process, n int, messages chan application.Message) {
	//input: the network# and the # of connections it will receive
	//listen to the client and decode the application, then send via channel
	var counter = 0
	ln, err := net.Listen("tcp", server.Ip + ":" + server.Port) //creates server
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		if counter == n{
			err = c.Close()
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		decoder := gob.NewDecoder(c) //initialize gob decoder

		//Decode application struct and print it
		mes := new(application.Message)
		_ = decoder.Decode(mes)

		t := time.Now()
		mes.T = t
		messages <- *mes
		counter++
	}
}
