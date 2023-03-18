/*package main

import (
	"bufio"
	"fmt"
	exit "guessWord"
	"net"
	"os"
	"strings"
)

func userInput(invitation string) string {
	fmt.Println(invitation)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("Ошибка при чтении ввода")
		return userInput(invitation)
	}
	return input
}

func getResponse(conn net.Conn) (string, error) {
	fmt.Println("я тут")
	reader := bufio.NewReader(conn)
	fmt.Println("и тут")
    input, err := reader.ReadBytes('\n')
	fmt.Println("и даже тут")
	if err != nil{
        return "", err
    }
	fmt.Println("и еще тут")
	return string(input), nil
}

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

	//userInputReader := bufio.NewReader(conn)
	//responseReader := bufio.NewReader(conn)

	ch := make(chan string)
	errCh := make(chan error)



	s := userInput("Введите число:")
	fmt.Println(s)
	fmt.Fprintf(conn, s)

	var response string
	var responseBytes []byte
	go func(ch chan string, errCh chan error) {
		for {
			conn.Read(responseBytes)
			response = string(responseBytes)
			//fmt.Fscan(conn, response)
			//response, err := getResponse(conn)
			fmt.Println("\"" + response + "\"")
            if err != nil {
                errCh <-err
                return
            }
            ch <-response
		}
	}(ch, errCh)

	for {
		select{
		case response := <-ch:
			if response == "userExit" {
				return
			}
			s = userInput("Введите число:")
			fmt.Fprintf(conn, s)
		case err := <-errCh:
			exit.EmergencyExit(1, err)
		}
	}
}*/