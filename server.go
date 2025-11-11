package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

type ReddishServer struct {
	port       int
	runStatus  int
	listener   net.Listener
	activeConn map[net.Conn]struct{} // Set of activeConn
}

func (r *ReddishServer) Start(port int) {
	defer r.Stop()

	r.activeConn = map[net.Conn]struct{}{}
	portStr := strconv.Itoa(port)
	addr := "localhost:" + portStr
	fmt.Println("Started listening on " + addr)

	// Get Listener
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Enocuntered error while listening " + err.Error())
		return
	}
	r.listener = listener
	r.runStatus = 1
	r.port = port

	for {
		conn, err := r.listener.Accept()
		if err != nil {
			fmt.Println("Encountered error while accepting" + err.Error())
			return
		}
		r.handleConnection(conn)
	}

}

func (r *ReddishServer) handleConnection(conn net.Conn) {
	defer r.closeConnection(conn)

	fmt.Printf("Accepted %s\n", conn.RemoteAddr().String())
	r.activeConn[conn] = struct{}{}

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

		fmt.Printf("Recieved: %v\n", string(buffer[:n]))

		ret := "+OK\r\n"
		_, err = conn.Write([]byte(ret))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}

func (r *ReddishServer) closeConnection(conn net.Conn) {
	fmt.Printf("Closed %s\n", conn.RemoteAddr().String())
	conn.Close()
	delete(r.activeConn, conn)
}

func (r *ReddishServer) Stop() {
	fmt.Println("Stopping services...")

	clear(r.activeConn)

	r.listener.Close()
	r.listener = nil

	r.port = 0
	r.runStatus = 0
}
