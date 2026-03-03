package utilidades

import (
	"github.com/gorilla/sessions"
	"net/http"
	gomail "gopkg.in/gomail.v2"
)

var Frontend string = "templates/layout/base.html"
var Store = sessions.NewCookieStore([]byte("session-name"))


func CrearMensajesFlash(response http.ResponseWriter, request *http.Request, css string, mensaje string) {
	session, err := Store.Get(request, "flash-session")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	session.AddFlash(css, "css")
	session.AddFlash(mensaje, "mensaje")
	session.Save(request, response)
}

func RetornarMensajesFlash(response http.ResponseWriter, request *http.Request) (string, string) {
	session, _ := Store.Get(request, "flash-session")

	fm := session.Flashes("css")
	session.Save(request, response)
	css_sesion := ""
	if len(fm) == 0 {
		css_sesion = ""
	} else {
		css_sesion = fm[0].(string)
	}
	fm2 := session.Flashes("mensaje")
	session.Save(request, response)
	css_mensaje := ""
	if len(fm2) == 0 {
		css_mensaje = ""
	} else {
		css_mensaje = fm2[0].(string)
	}
	return css_sesion, css_mensaje
}

func EnviarCorreo() {
	mensaje := gomail.NewMessage()
	mensaje.SetHeader("From", "")
	mensaje.SetHeader("To", "no.velizv@gmail.com")
	mensaje.SetHeader("Subject", "Aprendizaje Golang")
	mensaje.SetBody("text/html", "<h1>Aprendiendo Golang</h1><b>Este mensaje va en negritas</b><p>Este es un párrafo</p>")
	
	//Ahora se debe configurar la coneccion con el SMTP
	n:= gomail.NewDialer("smtp.dreamhost.com", 587, "correo","contaseña_correo")
	if err := n.DialAndSend(mensaje); err != nil {
		panic(err)
	}
}

func RetornarLogin(request *http.Request) (string, string) {
	session, _ := Store.Get(request, "session-name")
	hola_id := ""
	hola_nombre := ""
	if session.Values["hola_id"] != nil {
		hola_id_t, _ := session.Values["hola_id"].(string) // se crea una segunda variable para que no de error
		hola_id = hola_id_t
	}
		if session.Values["hola_nombre"] != nil {
		hola_nombre_t, _ := session.Values["hola_nombre"].(string) // se crea una segunda variable para que no de error
		hola_nombre = hola_nombre_t
	}
	return hola_id, hola_nombre

}