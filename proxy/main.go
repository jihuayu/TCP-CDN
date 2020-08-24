package main

import (
"os"
"fmt"
"net"
"golang.org/x/net/websocket"

)

func main()  {
	fmt.Print("You can join ",os.Args[1],". Connect ",os.Args[2],".")
	proxyStart(os.Args[1],os.Args[2])
}
func proxyStart(fromport, toport string) {
	proxylistener, err := net.Listen("tcp", fromport)
	if err != nil {
		fmt.Println("Unable to listen on: %s, error: %s\n", fromport, err.Error())
		os.Exit(1)
	}
	defer proxylistener.Close()

	for {
		proxyconn, err := proxylistener.Accept()
		if err != nil {
			fmt.Printf("Unable to accept a request, error: %s\n", err.Error())
			continue
		}

		buffer := make([]byte, 1024)
		n, err := proxyconn.Read(buffer)
		if err != nil {
			fmt.Printf("Unable to read from input, error: %s\n", err.Error())
			continue
		}


		targetconn, err := websocket.Dial(toport, "binary", "http://baidu.com");
		if err != nil {
			fmt.Println("Unable to connect to: %s, error: %s\n", toport, err.Error())
			proxyconn.Close()
			continue
		}

		n, err = targetconn.Write(buffer[:n])
		if err != nil {
			fmt.Printf("Unable to write to output, error: %s\n", err.Error())
			proxyconn.Close()
			targetconn.Close()
			continue
		}

		go proxyRequest(proxyconn, targetconn)
		go proxyRequest(targetconn, proxyconn)
	}
}

func proxyRequest(r net.Conn, w net.Conn) {
	defer r.Close()
	defer w.Close()

	var buffer = make([]byte, 4096000)
	for {
		n, err := r.Read(buffer)
		if err != nil {
			fmt.Printf("Unable to read from input, error: %s\n", err.Error())
			break
		}

		n, err = w.Write(buffer[:n])
		if err != nil {
			fmt.Printf("Unable to write to output, error: %s\n", err.Error())
			break
		}
	}
}