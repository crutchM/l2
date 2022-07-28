package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
	}
	con, _ := listener.Accept()

	for {
		msg, err := bufio.NewReader(con).ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Println("Server: ", msg)

		msg = "socket " + msg

		con.Write([]byte(msg + "\n"))
	}
}
