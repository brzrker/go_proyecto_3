package proteccion

import (
	"net/http"
	"proyecto_3/utilidades"
)

func Proteccion(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		session, _ := utilidades.Store.Get(request, "session-name")
		if session.Values["hola_id"] != nil {
			next.ServeHTTP(response, request)
		} else {
			utilidades.CrearMensajesFlash(response, request, "warning", "Debes estar logueado para ver este contenido")
			http.Redirect(response, request, "/seguridad/login", http.StatusSeeOther)
		}
	}
}


/* 
func Proteccion(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		next.ServeHTTP(response, request)
	}
}
*/