package main

import (
	"fmt"
	"io"
	"net"
)

const PORT string = "6379"
const IP string = "localhost"
const ADDRESS string = IP + ":" + PORT

func closeConnection(conn net.Conn) {
	fmt.Printf("Closed %s\n", conn.RemoteAddr().String())
	conn.Close()
}
func handleConnection(conn net.Conn) {
	defer closeConnection(conn)

	fmt.Printf("Accepted %s\n", conn.RemoteAddr().String())
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err == io.EOF {
			fmt.Println("Reached EOF")
			return
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		arg := string((buffer[:n]))
		fmt.Printf("Recieved:\n%s", arg)

		ret := "+OK\r\n"
		_, err = conn.Write([]byte(ret))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}
