# MP1

A simple network simulation implement with application layer and network layer

## To Run

There are two simulations in this project. \
The first simulation simulates **one process sends a message to another process.** \
The second simulation simulates **multiple processes send messages to a single process**.

### One to one

Open two terminals. In the first terminal, start a TCP server. In this case, it is the *process 1*.

```
go run main.go 1
```

Then in the second terminal, a user (process 2) wants to send message to the process 1.

```
go run main.go 2
```

Then you should see a guide look like this in your window:

```
please send application in this pattern: send 'processNumber' 'YourMessage'
```

So you follow the instruction:

```
send 1 hi
```

A notification should look like this:

```
Sent 'hi' to 1, system time is Sep 11 16:16:28.681
```

After a little bit delay, on your first terminal, you should receive a notification as well, indicates the message has been received:

```
Received 'hi' from 2, system time is Sep 11 16:16:30.687
```

### Multiple to One

It is similar to the one-to-one, except we need three terminals. One represents a server, two represent the clients.

In the first terminal, run following command:

```
go run main.go 3
```

In the second terminal, run:

```
go run main.go 2
```

send a message:

```
send 3 thisIs2
```

You should see something like this:

```
Sent 'thisIs2' to 3, system time is Sep 12 14:17:33.526
```

And the server will print:

```
Received 'thisIs2' from 2, system time is Sep 12 14:17:35.340
```

In the third terminal, run:
```
go run main.go 4
```

send a message:

```
send 3 thisIs4
```

You should see something like this:
```
Sent 'thisIs4' to 3, system time is Sep 12 14:19:16.937
```

In the server terminal you should see:

```
Received 'thisIs4' from 4, system time is Sep 12 14:19:19.763
```

Then the server will shut down.
## Structure and Design

### application

In `/application/message.go`, there are two struct:

```
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
```

The `Process` struct contains the basic information of a process, and the `Message` struct contains four elements: `S` represent source, `R` represent receiver, `M` represent the text being sent, `T` represent the time sent. \

It mainly simulates the application layer. The `GetInfo` function takes a process as input, instruct the user to send the message, and then returns a `Message` struct. The `UnicastReceive` function simply takes the source `Process` and `Message` as input, and print out the time received the message for the server.

### network

The `/network/clients.go` contains `UnicastSend` function, which takes the destination `Process` and `Message` as input, try to send the message to the destination via TCP connection. It is encoded with the support of `gob`. It also prints out the sent time and information for the client user.

The`/network/servers.go` contains `Server` function, which function as a go routine in the main function. It initialize a TCP server, and writes the message it receives to a go channel. The main function can then read the message from the channel and pass it to the `UnicastReceive` function. We also simulate the network delay in this function.

## Deployment
* [Channels and Go Routines](https://www.justindfuller.com/2020/01/go-things-i-love-channels-and-goroutines/)
* [Create a TCP and UDP Client and Server using Go](https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/)
* [Go Routines](https://golangbot.com/goroutines/)
* [MP0 by Jiahong Li and Zheng Zhou](https://github.com/jiahongli18/DistributedSystemsMP0)
* [Net Package](https://golang.org/pkg/net/)
* [Random Numbers](https://gobyexample.com/random-numbers)
* [Read Files](https://golangbot.com/read-files/)
* [Time Package](https://golang.org/pkg/time/)


## Authors

* **Gary Ge** - *initial work*
