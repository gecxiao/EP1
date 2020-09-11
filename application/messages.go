package application

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)
type Process struct {
	Id   string
	Ip   string
	Port string
}

type Message struct {
	S Process
	R string
	M string
	T time.Time
}

func GetInfo(client Process) Message {
	//get the application from user and pack in into Message struct
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please send application in this pattern: send 'processNumber' 'YourMessage'\n")
	text, _ := reader.ReadString('\n')
	t := strings.Fields(text)
	var m Message
	m.R = t[1]
	m.M = t[2]
	m.S = client
	return m
}

func UnicastReceive(source Process, message Message) {
	//delivers the application received from the source network.
	fmt.Printf("Received %s to %s, system time is %s\n", message.M, source.Id, message.T)
}