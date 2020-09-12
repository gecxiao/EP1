package main

/*
In this simulation
network 1 always wait a connection from network 2
network 3 always wait two connections: from network 2 and 4
The configuration file set max delay to be 3000(ms), min delay to be 1000(ms).
*/
import (
	"./application"
	"./network"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//open config file to load the processes
	data, err := ioutil.ReadFile("config.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	words := strings.Fields(string(data)) //[min_delay(ms) max_delay(ms) ID1 IP1 PORT1 ID2 IP2 PORT2 ...]
	var pA application.Process
	var pB application.Process
	var pC application.Process
	var pD application.Process
	minDelay, _ := strconv.Atoi(words[0])
	maxDelay, _ := strconv.Atoi(words[1])
	pA.Id, pA.Ip, pA.Port = words[2], words[3], words[4]
	pB.Id, pB.Ip, pB.Port = words[5], words[6], words[7]
	pC.Id, pC.Ip, pC.Port = words[8], words[9], words[10]
	pD.Id, pD.Ip, pD.Port = words[11], words[12], words[13]

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide the process number (1, 2, 3, 4).")
		return
	}
	pNum := arguments[1]

	switch {
	//Process 1 should receive a message from process 2.
	case pNum == "1":
		messages := make(chan application.Message)
		go network.Server(pA, 1, maxDelay, minDelay, messages)
		mes := <-messages
		application.UnicastReceive(mes.S, mes)
		return
	//Process 2 should send a message to either process 1 or process 3.
	case pNum == "2":
		mes := application.GetInfo(pB)
		if mes.R == "1" {
			network.UnicastSend(pA, mes)
			return
		} else if mes.R == "3" {
			network.UnicastSend(pC, mes)
			return
		} else {
			fmt.Println("Please send to process 1 or 3")
			return
		}
	//Process 3 will wait for two messages and then stop listening.
	case pNum == "3":
		messages := make(chan application.Message, 2)
		go network.Server(pC, 2, maxDelay, minDelay, messages)
		mes := <-messages
		application.UnicastReceive(mes.S, mes)
		mes = <-messages
		application.UnicastReceive(mes.S, mes)
		return
	//Process 4 should send a message to process 3.
	case pNum == "4":
		mes := application.GetInfo(pD)
		if mes.R == "3" {
			network.UnicastSend(pC, mes)
			return
		} else {
			fmt.Println("Please send to process 3")
			return
		}
	default:
		fmt.Println("Please provide the network number (1, 2, 3, 4).")
		return
	}
}
