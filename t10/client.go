package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	timeOut := flag.Int("timeout", 10, "time out of connection")
	flag.Parse()

	if len(os.Args) < 4 {
		log.Fatal("Not enough arguments.\nUSE : -timeout <timeout> <host> <port>")
	}

	con, err := net.DialTimeout("tcp", os.Args[3]+":"+os.Args[4], time.Duration(*timeOut)*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	if con != nil {
		defer con.Close()
		log.Println("Connection is opened")
	}

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err == io.EOF {
				con.Close()
			}
			fmt.Fprintf(con, text+"\n")
		}
	}()

	for {
		msg, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("Client: " + msg)
	}
}
