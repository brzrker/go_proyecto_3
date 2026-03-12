package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proyecto_3/proteccion"
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
	mux.HandleFunc("/estructuras", rutas.Estructuras)

	// Formulario
	mux.HandleFunc("/formulario", rutas.Formulario_get)
	mux.HandleFunc("/formulario-post", rutas.Formulario_post).Methods("POST")
	mux.HandleFunc("/formulario/upload", rutas.Formulario_upload)
	mux.HandleFunc("/formulario/upload-post", rutas.Formulario_upload_post).Methods("POST")

	// Recursos utiles
	mux.HandleFunc("/recursos-utiles", proteccion.Proteccion(rutas.Recursos_utiles_get))
	mux.HandleFunc("/recursos-utiles/pdf", proteccion.Proteccion(rutas.Recursos_utiles_pdf))
	mux.HandleFunc("/recursos-utiles/pdf-generar", proteccion.Proteccion(rutas.Recursos_utiles_pdf_generar))
	mux.HandleFunc("/recursos-utiles/excel", proteccion.Proteccion(rutas.Recursos_utiles_excel))
	mux.HandleFunc("/recursos-utiles/qr", proteccion.Proteccion(rutas.Recursos_utiles_qr))
	mux.HandleFunc("/recursos-utiles/enviar-correo", proteccion.Proteccion(rutas.Recursos_utiles_enviar_correo))

	// Cliente http
	mux.HandleFunc("/cliente-http", proteccion.Proteccion(rutas.Cliente_http))
	mux.HandleFunc("/cliente-http/cliente-http-crear", proteccion.Proteccion(rutas.Cliente_http_crear))
	mux.HandleFunc("/cliente-http/cliente-http-crear-post", proteccion.Proteccion(rutas.Cliente_http_crear_post)).Methods("POST")
	mux.HandleFunc("/cliente-http/cliente-http-editar/{id:.*}", proteccion.Proteccion(rutas.Cliente_http_editar))
	mux.HandleFunc("/cliente-http/cliente-http-editar-post/{id:.*}", proteccion.Proteccion(rutas.Cliente_http_editar_post)).Methods("POST")
	mux.HandleFunc("/cliente-http/cliente-http-eliminar/{id:.*}", proteccion.Proteccion(rutas.Cliente_http_eliminar))

	// MySQL
	mux.HandleFunc("/my-sql", proteccion.Proteccion(rutas.My_SQL_listar))
	mux.HandleFunc("/my-sql/my-sql-crear", proteccion.Proteccion(rutas.My_SQL_crear))
	mux.HandleFunc("/my-sql/my-sql-crear-post", proteccion.Proteccion(rutas.My_SQL_crear_post)).Methods("POST")
	mux.HandleFunc("/my-sql/my-sql-editar/{id:.*}", proteccion.Proteccion(rutas.My_SQL_editar))
	mux.HandleFunc("/my-sql/my-sql-editar-post/{id:.*}", proteccion.Proteccion(rutas.My_SQL_editar_post)).Methods("POST")
	mux.HandleFunc("/my-sql/my-sql-eliminar/{id:.*}", proteccion.Proteccion(rutas.My_SQL_eliminar))

	// Seguridad
	mux.HandleFunc("/seguridad/registro", rutas.Seguridad_registro)
	mux.HandleFunc("/seguridad/registro-post", rutas.Seguridad_registro_post).Methods("POST")
	mux.HandleFunc("/seguridad/login", rutas.Seguridad_login)
	mux.HandleFunc("/seguridad/login-post", rutas.Seguridad_login_post).Methods("POST")
	mux.HandleFunc("/seguridad/protegida", proteccion.Proteccion(rutas.Seguridad_protegida))
	mux.HandleFunc("/seguridad/logout", proteccion.Proteccion(rutas.Seguridad_logout))

	// Pasarelas
	mux.HandleFunc("/pasarelas", proteccion.Proteccion(rutas.Pasarelas))
	mux.HandleFunc("/pasarelas/webpay", proteccion.Proteccion(rutas.Pasarelas_webpay))
	mux.HandleFunc("/pasarelas/webpay/retorno", proteccion.Proteccion(rutas.Pasarelas_webpay_retorno))
	mux.HandleFunc("/pasarelas/paypal", proteccion.Proteccion(rutas.Pasarelas_paypal))
	mux.HandleFunc("/pasarelas/paypal/retorno", proteccion.Proteccion(rutas.Pasarelas_paypal_retorno))
	mux.HandleFunc("/pasarelas/paypal/cancelar", proteccion.Proteccion(rutas.Pasarelas_paypal_cancelar))

	// Archivos estaticos
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	mux.PathPrefix("/static/").Handler(s)

	// Manejo de errores
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(rutas.Pagina404).GetHandler()

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