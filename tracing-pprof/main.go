package main

import (
	"fmt"
	"math/rand"
	"time"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	// Importando o pacote net/http/pprof para habilitar os perfis do pprof
	go func() {
		fmt.Println("Starting pprof server on :6060")
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			fmt.Println("Error starting pprof server:", err)
		}
	}()

	for i := 0; i < 10; i++ {
		go calculation()
	}

	// Para manter o programa rodando para coletar o perfil do pprof
	fmt.Println("Pressione Ctrl+C para encerrar.")
	select {}
}

func calculation() {
	for {
		result := 0
		for i := 0; i < 100000; i++ {
			// Realizando algum cálculo intensivo de CPU
			result += rand.Intn(1000)
		}
		fmt.Println("Result:", result)
		time.Sleep(100 * time.Millisecond)
	}
}
