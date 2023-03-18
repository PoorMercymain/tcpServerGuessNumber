/*package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	exit "guessWord"
)

func main() {

	port := "8080"

	builder := strings.Builder{}

	builder.WriteString("localhost:")
	builder.WriteString(port)

    conn, err := net.Dial("tcp", builder.String())
    if err != nil {
		exit.EmergencyExit(1, err)
    }
    defer conn.Close()

	reader := bufio.NewReader(conn)
	var response string
	ch := make(chan string)
	eCh := make(chan error)

	go func(ch chan string, eCh chan error) {
  		for {
    		response, err = reader.ReadString('\n')
    		if err != nil {
      			eCh<- err
      			return
    		}
    		ch<- response
  		}
	}(ch, eCh)

	userInputScanner := bufio.NewScanner(os.Stdin)
	var firstTime bool = true
	var text string

	if firstTime {
		userInputScanner.Scan()
		text = userInputScanner.Text()
		fmt.Fprintln(conn, text)
	}

	for {
  		select {
     	case response = <-ch:
			fmt.Println("\"" + response + "\"")
			if response == "quit" {
				fmt.Println("Соединение разорвано со стороны сервера")
				return
			} else if response == "quit from client" {
				fmt.Println("Соединение прервано вами")
				return
			}

			fmt.Println("тего")
			userInputScanner.Scan()
			text := userInputScanner.Text()
			fmt.Println("\"" + text + "\"")
			fmt.Fprintf(conn, text)
     	case err = <-eCh:
			fmt.Println("f")
			exit.EmergencyExit(1, err)
	 	}
	}
}*/