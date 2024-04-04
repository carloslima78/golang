package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Handlers para cada verbo REST
	http.HandleFunc("/", handleDefault)
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/post", handlePost)
	http.HandleFunc("/put", handlePut)
	http.HandleFunc("/delete", handleDelete)

	// Rodando o servidor na porta 8082
	port := ":8082"
	fmt.Printf("Servidor rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo à API REST Simples!")
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET: Recebido!")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST: Recebido!")
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PUT: Recebido!")
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE: Recebido!")
}
