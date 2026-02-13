package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	//mux := http.NewServeMux()

	//mux.HandleFunc("/", func(response http.ResponseWriter, request * http.Request){
	http.HandleFunc("/", func(response http.ResponseWriter, request * http.Request){ // Lo mismo pero sin usar el servidor mux
		fmt.Fprintln(response, "hola mundo")
	})
	fmt.Println("corriendo servidor desde http://192.168.1.88:8081")
	log.Fatal(http.ListenAndServe("192.168.1.88:8081", nil))
}