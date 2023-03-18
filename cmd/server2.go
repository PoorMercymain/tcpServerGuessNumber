/*package main

import (
	"fmt"
	exit "guessWord"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)



func handleConnection(conn net.Conn) {
	defer conn.Close()

	//reader := bufio.NewReader(conn)
	//writer := bufio.NewWriter(conn)

	//readWriter := bufio.NewReadWriter(reader, writer)

	number := rand.Intn(100)+1
	fmt.Println(number)
	var userInputBytes []byte
	for {
		_, err := conn.Read(userInputBytes)
		if err != nil {
			fmt.Println(err)
            return
		}
		fmt.Println(string(userInputBytes))
        if string(userInputBytes) == "exit" {
        	fmt.Println("userExit")
			break
        }
		fmt.Println("\"" + strconv.Itoa(number) + "\"")
		numberStr := strconv.Itoa(number)
		numberBytes := []byte(numberStr)
		conn.Write(numberBytes)
		//writer.WriteString(strconv.Itoa(number))

		fmt.Println("отправил", number)
	}
}

func main() {
	port := "8080"
	rand.Seed(time.Now().UnixNano())
	builder := strings.Builder{}

	builder.WriteString(":")
	builder.WriteString(port)

    l, err := net.Listen("tcp", builder.String())
	if err != nil {
		exit.EmergencyExit(1, "Ошибка при попытке начать прослушивание порта", port, "-", err)
    }
    defer l.Close()

	fmt.Println("Сервер на", l.Addr().String(), "запущен и ожидает подключений...")

	for {
		conn, err := l.Accept()
        if err!= nil {
			fmt.Fprintln(os.Stderr, "Ошибка при попытке принять подключение -", err)
            continue
        }

		fmt.Println("Подключение", conn.RemoteAddr())

        go handleConnection(conn)
	}
}*/