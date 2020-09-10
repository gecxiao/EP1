package main

import (
	"./message"
	"fmt"
	"io/ioutil"
	"strings"
)
/*
Configuration file
min_delay(ms) max_delay(ms)
ID1 IP1 port1
ID2 IP2 port2

 */
func main(){
	//open config file to load the processes
	data, err := ioutil.ReadFile("config.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	words := strings.Fields(string(data)) //[min_delay(ms) max_delay(ms) 1 IP1 8080 2 IP2 8081 ...]
	var clientA message.Address
	var clientB message.Address
	var serverA message.Address
	var serverB message.Address
	minDelay := words[0]
	maxDelay := words[1]
	clientA.Id, clientA.Ip, clientA.Port = words[2], words[3], words[4]
	clientB.Id, clientB.Ip, clientB.Port = words[5], words[6], words[7]
	serverA.Id, serverA.Ip, serverA.Port = words[8], words[9], words[10]
	serverB.Id, serverB.Ip, serverB.Port = words[11], words[12], words[13]


	//getInfo(client) --> msg, sender, receiver
	//go message.unicast_send()
	//go server --> message (type message)
	//unicast_receive(source, message)  print the message?
}