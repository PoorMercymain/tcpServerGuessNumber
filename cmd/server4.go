/*package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	exit "guessWord"
)

func closeMultipleConnections(connections []net.Conn) {
	for _, conn := range connections {
		fmt.Fprintf(conn, "quit")
		if err := conn.Close(); err != nil {
			continue
		}
    }
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("что я здесь делаю...")

	scanner := bufio.NewScanner(conn)
	var fromClient string

	number := rand.Intn(100) + 1

	var counter int

	for scanner.Scan() {
		counter++
		fromClient = scanner.Text()
		fmt.Println(fromClient)

		if fromClient == "quit" {
			fmt.Println(fromClient)
			fmt.Fprintf(conn, "quit from client")
            return
		}

    	fmt.Fprintf(conn, "Случайное число:", number)
		fmt.Println("впвпварппр", number)

		fmt.Println("я получил:\"" + fromClient + "\"", "на", counter, "раз")
	}
}

func main() {
	connections := make([]net.Conn, 0)
	defer closeMultipleConnections(connections)

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

    fmt.Println("Сервер запущен и ожидает подключений...")

    for {
        conn, err := l.Accept()
		if err != nil {
			fmt.Println("Ошибка при попытке принять подключение -", err)
            continue
		}

        fmt.Fprintln(os.Stdout, "Подключение от", conn.RemoteAddr())

		connections = append(connections, conn)

        go handleConnection(conn)
    }
}*/