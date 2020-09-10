package process

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
	"../message"
)

func server(server message.Address) message.Message {
	//handle a connection, waiting for the dialer,
	//should return the message received and the time received.
	ln, err := net.Listen("tcp", server.Ip) //creates server
	if err != nil {
		fmt.Println(err) // handle error
	}
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err) // handle error
	}
	m, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	now := time.Now()
	err = conn.Close()
	if err != nil {
		log.Panic(err)
	}
	var mes message.Message
	mes.M = m
	mes.T = now
	return mes
}
