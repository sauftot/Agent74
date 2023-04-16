package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/netip"
	"os"
	"strings"
	"time"
)

var run uint8 = 1

func readInput(ch chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		ch <- text
	}
}

//goroutines

func A74socket() {
	port := "24765"
}

func extSocket() {
	port := "25022"
}

func controllerSocket(ctrCh <-chan string) {
	// Define the address and port to listen on
	var port uint16
	port = 24666
	var connected = false
	expectedData := ""
	netTCPAddr := net.TCPAddrFromAddrPort(netip.AddrPortFrom(netip.IPv4Unspecified(), port))
	listener, err := net.ListenTCP("tcp", netTCPAddr)
	defer func(listener *net.TCPListener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("ERROR: controllerSocket | listener.Close()")
		}
	}(listener)
	if err != nil {
		fmt.Println("ERROR: controllerSocket | net.Listen()")
	}

	// Create a TCP listener on the specified address and port
	for run == 1 {
		// Wait for a client to connect
		var conn net.Conn
		for run == 1 {
			err := listener.SetDeadline(time.Now().Local().Add(time.Second))
			if err != nil {
				fmt.Println("ERROR: controllerSocket | listener.SetDeadline()")
				return
			}
			conn, err = listener.Accept()
			if err != nil {
				fmt.Println("ERROR: controllerSocket | listener.Accept()")
				return
			} else {
				break
			}
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				fmt.Println("ERROR: controllerSocket | defer conn.Close()")
			}
		}(conn)

		reader := bufio.NewReader(conn)
		writer := bufio.NewWriter(conn)
		var n int
		for n == len(expectedData) {
			n = reader.Buffered()
		}

		receivedData, _ := reader.ReadString('\n')
		if strings.TrimRight(receivedData, "\n") == strings.TrimRight(expectedData, "\n") {
			connected = true
			_, err := writer.WriteString("ok")
			if err != nil {
				fmt.Sprintln()
			}
			writer.Flush()
			fmt.Println("INFO: A74Client connected")
			//authentication done
			for connected && run == 1 {
				//read from control channel to handle case of application termination
				select {
				case i := <-ctrCh:
					if strings.Contains(i, "open") {
						//tell A74client to open connection
						writer.WriteString("open")
						writer.Flush()
					}
				}

				//read from tcp connection for
				n = reader.Buffered()
				if n > 0 {
					_, err1 := reader.ReadString('\n')
					if err1 != nil {
						if err1 == io.EOF {
							connected = false
							break
						} else {
							fmt.Println("ERROR: controllerSocket | read from A74 client | error not EOF")
						}
					}
				}
			}
		}

	}

	// Close the connection and listener
	return
}

func controller(ch <-chan string, ctr chan<- int) {

}

func main() {
	//input from std
	inputCh := make(chan string)
	//output to std
	outputCh := make(chan string)
	//ctrl to ctrlSocket
	ctrCh := make(chan string)
	//ctrlSocket to ctrl
	ctrSockCh := make(chan int)
	slaveCh := make(chan int)

}
