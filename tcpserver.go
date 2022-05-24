package main

import (
	"bufio"
	"fmt"
	"net"
)

//See: https://www.youtube.com/watch?v=l3KCOjhFXbQ
//Test using PuTTY telnet connection to listening port

func main() {
	//Listen
	//Accept Connection
	//Handle Connection -> Thread

	dstream, err := net.Listen("tcp", ":5044")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dstream.Close()

	for {
		conn, err := dstream.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		//go routine
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	}
	conn.Close()
}
