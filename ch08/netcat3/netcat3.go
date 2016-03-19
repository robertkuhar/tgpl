package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	doneChannel := make(chan struct{})
	go func() {
		// NOTE: ignoring errors
		io.Copy(os.Stdout, conn)
		log.Println("done")
		// signal the main goroutine
		doneChannel <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-doneChannel
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst,src); err != nil {
		log.Fatal(err)
	}
}

