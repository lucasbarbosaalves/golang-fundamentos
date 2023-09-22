package main

import (
	"log"
	"net/http"
)

/*
 HTTP é um protocolo de comunicação - base da comunicação web
 É a ponte de comunicação entre cliente e servidor

 Cliente faz uma requisição para um servidor e o servidor responde
 A resposta pode ser um HTML, JSON, XML, arquivo, etc
*/

func main() {

	// Http -> pacote nativo para trabalhar com requisicao http
	// ListenAndServe -> sobe um servidor web
	// Primeiro parametro é a porta e o segundo é um handler

	http.HandleFunc("/home", func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Olá Mundo!")) // Escreve na resposta um byte de slice
	})

	http.HandleFunc("/usuarios", func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("Carregar usuarios!"))
	})
	
	
	log.Fatal(http.ListenAndServe(":5000", nil))

}