package message

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)
type Address struct {
	Id   string
	Ip   string
	Port string
}

type Message struct {
	M string
	T time.Time
}

func getInfo(client Address) (string, string, string) {
	//get the message
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please send message in this pattern: send 'processNumber' 'YourMessage'")
	text, _ := reader.ReadString('\n')
	t := strings.Fields(string(text))
	receiver := t[1]
	message := t[2]
	sender := client.Id

	return message, receiver, sender
}

func unicast_receive(source Address, message Message) {
	//delivers the message received from the source process.
	fmt.Printf("Received %s to %s, system time is %s\n", message.M, source.Id, message.T)
}