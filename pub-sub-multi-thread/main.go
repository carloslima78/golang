package main

import "fmt"

func main() {
	ch := make(chan string)
	go publish(ch)
	consume(ch)
}

func publish(ch chan string) {
	mensagens := []string{"Olá", "Mundo"}

	for _, mensagem := range mensagens {
		ch <- mensagem
		fmt.Println("Mensagem publicada:", mensagem)
	}
	close(ch)
}

func consume(ch chan string) {
	for mensagem := range ch {
		fmt.Println("Mensagem consumida:", mensagem)
	}
}
