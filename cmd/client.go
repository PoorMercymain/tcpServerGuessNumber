package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	serverAddr := "localhost:8080"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err.Error())
		return
	}
	defer conn.Close()

	for {
		fmt.Print("Введите 'exit' для выхода или число: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при чтении ввода:", err.Error())
			continue
		}

		input = strings.TrimSpace(input)

		for len([]byte(input)) > 300 {
			fmt.Println("Ошибка при чтении ввода: слишком длинный ввод")
			input, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка при чтении ввода:", err)
			}
		}
		if _, err := conn.Write([]byte(input + "\n")); err != nil {
			fmt.Println("Ошибка при отправке запроса:", err.Error())
			continue
		}

		if input == "exit" {
			break
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении ответа от сервера:", err.Error())
			continue
		}

		response := strings.TrimSpace(string(buffer[:n]))
		if response == "" {
			fmt.Println("Получен пустой ответ от сервера")
			continue
		}

		fmt.Println("Получен ответ:", response)
		if response == "EQUAL" {
			break
		}
	}

	fmt.Println("Отключение от сервера...")
}
