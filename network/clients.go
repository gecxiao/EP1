package network

import (
	"../application"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func UnicastSend(destination application.Process, m application.Message) {
	//Sends application to the destination network.
	//simulate network delay here.
	c, err := net.Dial("tcp", destination.Ip+":"+destination.Port) //connect to TCP server
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	fmt.Printf("Sent '%s' to %s, system time is %s\n", m.M, destination.Id, now)
	r := rand.Float64() * 3 + 1
	time.Sleep(time.Duration(r) * time.Second)
	encoder := gob.NewEncoder(c)
	msg := application.Message{
		S: m.S,
		R: m.R,
		M: m.M,
		T: time.Time{},
	}
	_ = encoder.Encode(msg)
}
