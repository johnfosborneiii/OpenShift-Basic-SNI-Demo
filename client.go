package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	hostname := os.Args[1] // add some error checking
	server := flag.String("server", hostname + ":443", "the server and port to connect to in the form of 10.0.2.15:443")
	flag.Parse()

	config := &tls.Config{
		ServerName:         hostname,
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", *server, config)
	if err != nil {
		panic(err)
	}
	err = conn.Handshake()
	if err != nil {
		fmt.Printf("Failed handshake:%v\n", err)
		return
	}

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Printf("Failed receiving data:%v\n", err)
	}

	conn.Close()
}
