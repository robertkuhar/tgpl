package main

import (
	"time"
	"net"
	"fmt"
	"strings"
	"log"
	"bufio"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// BobK:  go on the handleConn gets us multiple, simultaneous connections!
		go handleConn(conn)
	}

}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		// echo(c, input.Text(), 1 * time.Second)
		go echo(c, input.Text(), 1 * time.Second)
	}
	log.Println("Imma gonna c.Close(), are the echos done?")
	// NOTE: ignoring potential errors from input.Err
	c.Close()
}
