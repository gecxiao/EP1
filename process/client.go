package process

import (
	"fmt"
	"log"
	"net"
	"time"
	"../message"
)

func unicast_send(destination message.Address, message string) {
	//Sends message to the destination process.
	c, err := net.Dial("tcp", destination.Ip) //connect to TCP server
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(c, message+"\n")

	now := time.Now()
	fmt.Printf("Sent %s to %s, system time is %s\n", message, destination.Id, now)
}
