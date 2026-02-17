package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"proyecto_3/rutas"
	"time"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	mux := mux.NewRouter()
	//Rutas
	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/nosotros", rutas.Nosotros)
	mux.HandleFunc("/parametros/{id:.*}/{slug:.*}", rutas.Parametros)
	mux.HandleFunc("/parametros-querystring", rutas.ParametrosQueryString)
	// Ejecucion de servidor
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}

	server := &http.Server{
		Addr: "localhost:"+ os.Getenv("PORT"),
		Handler: mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	fmt.Println("corriendo servidor desde http://localhost:"+ os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
}

/* func main() {
	//mux := http.NewServeMux()

	//mux.HandleFunc("/", func(response http.ResponseWriter, request * http.Request){
	http.HandleFunc("/", func(response http.ResponseWriter, request * http.Request){ // Lo mismo pero sin usar el servidor mux
		fmt.Fprintln(response, "hola mundo")
	})
	fmt.Println("corriendo servidor desde http://192.168.1.88:8081")
	log.Fatal(http.ListenAndServe("192.168.1.88:8081", nil))
} */