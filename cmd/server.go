package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100) + 1

	fmt.Printf("Получен запрос на угадывание числа от %s\n", conn.RemoteAddr().String())
	fmt.Println("Загаданное число:", number)

	for {
		buffer := make([]byte, 1024)
		readLen, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении полученных от клиента данных:", err.Error())
			return
		}

		inputStr := strings.TrimSpace(string(buffer[:readLen]))
		if inputStr == "exit" {
			fmt.Println("Соединение закрыто клиентом")
			return
		}

		var response string

		input := strings.Split(inputStr, " ")
		fmt.Println(input)
		if len(input) < 2 {
			fmt.Println("Ошибка: команда на угадывание числа должна содержать команду и число")
			response = "Команда на угадывание числа должна содержать команду и число"
			_, err = conn.Write([]byte(response))
			if err != nil {
				fmt.Println("Ошибка при отправке ответа:", err.Error())
				return
			}
			fmt.Printf("Информация об ошибке отправлена клиенту %s\n", conn.RemoteAddr().String())
			continue
		}

		if input[0] != "GUESS" {
			fmt.Println("Некорректная команда: ожидалось GUESS")
			response = "Некорректная команда"
			_, err = conn.Write([]byte(response))
			if err != nil {
				fmt.Println("Ошибка при отправке ответа:", err.Error())
				return
			}
			fmt.Printf("Информация об ошибке отправлена клиенту %s\n", conn.RemoteAddr().String())
			continue
		}

		inputNumber, err := strconv.Atoi(input[1])
		if err != nil {
			fmt.Println("\"" + inputStr + "\"")
			fmt.Println("Ошибка при конвертации запроса в число:", err)
			response = "Ошибка конвертации"
			_, err = conn.Write([]byte(response))
			if err != nil {
				fmt.Println("Ошибка при отправке ответа:", err.Error())
				return
			}
			fmt.Printf("Информация об ошибке отправлена клиенту %s\n", conn.RemoteAddr().String())
			continue
		}

		if number < inputNumber {
			response = "LESS"
		} else if number > inputNumber {
			response = "MORE"
		} else {
			response = "EQUAL"
		}

		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Ошибка при отправке ответа:", err.Error())
			return
		}
		fmt.Printf("Отправлен ответ %s клиенту %s\n", response, conn.RemoteAddr().String())
		if response == "EQUAL" {
			fmt.Println("Соединение закрыто")
			return
		}
	}
}

func main() {
	port := "8080"
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Ошибка при попытке начать прослушивание:", err.Error())
		return
	}
	defer l.Close()
	fmt.Println("Сервер слушает на порте " + port + "...")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии подключения:", err.Error())
			continue
		}
		fmt.Println("Принято подключение от:", conn.RemoteAddr().String())

		go handleConnection(conn)
	}
}
