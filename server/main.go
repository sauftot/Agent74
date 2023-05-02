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
	"math/rand"
)

type A74Conn struct {
	external *net.TCPConn
	a74 *net.TCPConn
	id uint64
}

var run uint8 = 1
var connected uint8 = 0


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
	//port := "24765"
}

func extSocket() {
	//port := ""
	var port uint16
	port = 25008
	netTCPAddr := net.TCPAddrFromAddrPort(netip.AddrPortFrom(netip.IPv4Unspecified(), port))
	listener, err := net.ListenTCP("tcp", netTCPAddr)
	if err != nil {
		fmt.Println("ERROR: extSocket | net.ListenTCP()")
	}
	defer func(listener *net.TCPListener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("ERROR: controllerSocket | listener.Close()")
		}
	}(listener)

	var conn *net.TCPConn
	for run == 1 {
		for connected == 1 && run == 1 {
			err := listener.SetDeadline(time.Now().Local().Add(time.Millisecond * 20))
			if err != nil {
				fmt.Println("ERROR: controllerSocket | listener.SetDeadline()")
				return
			}
			conn, err = listener.AcceptTCP()
			if err != nil {
				fmt.Println("ERROR: controllerSocket | listener.Accept()")
				return
			} else {
				conn := A74Conn{external: conn, id: rand.Uint64()}
				
				go extConn()
			}
		}

	}



}

func extConn(A74Conn) {

}

func controllerSocket(ctrlToExtSocks chan<- A74Conn, extToCtrlSocks <-chan A74Conn) {
	// Define the address and port to listen on
	var port uint16
	port = 24666
	expectedData := ""
	netTCPAddr := net.TCPAddrFromAddrPort(netip.AddrPortFrom(netip.IPv4Unspecified(), port))
	listener, err := net.ListenTCP("tcp", netTCPAddr)
	if err != nil {
		fmt.Println("ERROR: controllerSocket | net.ListenTCP()")
	}
	defer func(listener *net.TCPListener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("ERROR: controllerSocket | listener.Close()")
		}
	}(listener)

	var conn *net.TCPConn
	_ = conn.SetKeepAlive(true)
	_ = conn.SetKeepAlivePeriod(time.Second)


	// Create a TCP listener on the specified address and port
	for run == 1 {
		// Wait for a client to connect
		for run == 1 {
			err := listener.SetDeadline(time.Now().Local().Add(time.Second))
			if err != nil {
				fmt.Println("ERROR: controllerSocket | listener.SetDeadline()")
				return
			}
			conn, err = listener.AcceptTCP()
			if err != nil {
				fmt.Println("ERROR: controllerSocket | listener.Accept()")
				return
			} else {
				break
			}
		}
		defer func(conn *net.TCPConn) {
			err := conn.Close()
			if err != nil {

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
			connected = 1
			_, err := writer.WriteString("ok")
			if err != nil {
				fmt.Sprintln()
			}
			writer.Flush()
			fmt.Println("INFO: A74Client connected")
			//authentication done
			sh := make([]byte, 1024)
			for connected == 1 && run == 1 {
				//read from extToCtrlSocks channel, if there is a new A74Conn struct in there tell the controller to open a tcp connection,
				//listen for that tcp connection then handle the entire A74Conn struct 
				select {
				case

				}

				//check if connection is alive
				_, err := conn.Read(sh) //BLOCKING!!!!!!!
				if err != nil {
					if err == io.EOF {
						connected = 0
						reader = nil
						writer = nil
						break
					} else {
						fmt.Println("ERROR: controllerSocket | error keeping connection alive")
					}
				}
			}
		}

	}

	// Close the connection and listener
	return
}

func controller(ch <-chan string) {
	var cmd string
	for {
		cmd = <-ch
		if strings.Contains(strings.ToLower(cmd), "stop") {
			run = 0
			connected = 0
			return
		}
	}
}

func main() {
	//input from std
	/*inputCh := make(chan string)
	//output to std
	outputCh := make(chan string)
	//ctrl to ctrlSocket
	ctrCh := make(chan string)
	//ctrlSocket to ctrl
	ctrSockCh := make(chan int)
	slaveCh := make(chan int)
	*/

	fmt.Println("hi")
}
