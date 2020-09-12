package network

import (
	"../application"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"time"
)

func UnicastSend(destination application.Process, m application.Message) {
	//Sends message to the destination process.
	c, err := net.Dial("tcp", destination.Ip+":"+destination.Port) //connect to TCP server
	if err != nil {
		log.Panic(err)
	}
	now := time.Now().Format("Jan _2 15:04:05.000")
	fmt.Printf("Sent '%s' to %s, system time is %s\n", m.M, destination.Id, now)
	encoder := gob.NewEncoder(c)
	msg := application.Message{
		S: m.S,
		R: m.R,
		M: m.M,
		T: now,
	}
	_ = encoder.Encode(msg)
	return
}
