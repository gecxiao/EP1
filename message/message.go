package message

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type Address struct {
	id   string
	ip   string
	port string
}

type Info struct {
	Message string
	Time    time.Time
}

func Listener(server Address) (string, time.Time) {
	//handle a connection, waiting for the dialer,
	//should return the message received and the time received.
	ln, err := net.Listen("tcp", server.ip) //creates server
	if err != nil {
		fmt.Println(err) // handle error
	}
	for {
		conn, err := ln.Accept()
		now := time.Now()
		if err != nil {
			fmt.Println(err) // handle error
		}
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		err = conn.Close()
		if err != nil{
			log.Fatal(err)
		}
		return message, now
	}
}

func Dialer(client Address) (string, time.Time){
	//try to build a connection to the server,
	//encode the message and send to the server,
	//return message and the time sent.
	c, err := net.Dial("tcp", client.ip) //connect to TCP server
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin) //read the message from the user and send it to the server
	fmt.Print("please send message in this pattern: send 'processNumber' 'YourMessage'")
	message, _ := reader.ReadString('\n')
	fmt.Fprintf(c, message+"\n")

	now := time.Now()
	return message, now
}

func unicast_send(destination Address, message string) {
	//Sends message to the destination process.
	//a little bit confused (because the message should be entered by the user,
	//but here it appears as an argument in the function.
	message, t := Dialer(destination)
	fmt.Printf("Sent %s to %s, system time is %s\n", message, destination.id, t)
}

func unicast_receive(source Address, message string) {
	//delivers the message received from the source process.
	message, t := Listener(source)
	fmt.Printf("Sent %s to %s, system time is %s\n", message, destination.id, t)
}
