package rutas

import (
	"net/http"
	"fmt"
)

func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "hola mundo desde golang")
}