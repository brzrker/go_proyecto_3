package rutas

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"proyecto_3/modelos"
	"proyecto_3/utilidades"
	"strconv"
//	"strings"
//	"time"

	"github.com/joho/godotenv"
	paypal "github.com/plutov/paypal/v4"
)

func Pasarelas(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/pasarelas/pasarelas.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css":     css_sesion,
		"mensaje": css_mensaje,
	}
	template.Execute(response, data)
}

func Pasarelas_webpay(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/pasarelas/webpay.html", utilidades.Frontend))
	// Comunicacion a webpay
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		fmt.Println(errorVariables)
		return
	}
	// Instancia Cliente
	cliente := &http.Client{}
	// Datos
	datos := map[string]string{"buy_order":"ordenCompra123123123", "session_id":"orderCompra123123123", "amount":"10000", "return_url":"http://localhost:8080/pasarelas/webpay/retorno"}
	jsonValue, _ := json.Marshal(datos)
	req, err := http.NewRequest("POST", os.Getenv("WEBPAY_URL"), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Tbk-Api-Key-Id", os.Getenv("WEBPAY_ID"))
	req.Header.Set("Tbk-Api-Key-Secret", os.Getenv("WEBPAY_SECRET"))
	if err != nil {
		fmt.Println(err)
		return
	}
	reg, err2 := cliente.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer reg.Body.Close()

	body, err := io.ReadAll(reg.Body)
	webpay := modelos.WebpayModel{}
	errjson := json.Unmarshal(body, &webpay)
	if errjson != nil {
		fmt.Println(errjson)
		return
	}
	// Retorno
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css": css_sesion,
		"mensaje": css_mensaje,
		"url": webpay.Url,
		"token": webpay.Token,
	}
	template.Execute(response, data)
}

func Pasarelas_webpay_retorno(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/pasarelas/webpay_retorno.html", utilidades.Frontend))

	errorVariables := godotenv.Load()
	if errorVariables != nil {
		fmt.Println(errorVariables)
		return
	}
	cliente := &http.Client{}
	req, err := http.NewRequest("PUT", os.Getenv("WEBPAY_URL")+"/"+request.URL.Query().Get("token_ws"), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Tbk-Api-Key-Id", os.Getenv("WEBPAY_ID"))
	req.Header.Set("Tbk-Api-Key-Secret", os.Getenv("WEBPAY_SECRET"))
	if err != nil {
		fmt.Println(err)
		return
	}
	reg, err2 := cliente.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer reg.Body.Close()

	body, err := io.ReadAll(reg.Body)
	webpay := modelos.WebpayRetornoModel{}
	errjson := json.Unmarshal(body, &webpay)
	if errjson != nil {
		fmt.Println(errjson)
		return
	}

	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	amount := strconv.Itoa(webpay.Amount)
	response_code := strconv.Itoa(webpay.Response_code)
	installments_number := strconv.Itoa(webpay.Installments_number)
	data := map[string]string{
		"css": css_sesion,
		"mensaje": css_mensaje,
		"token_ws": request.URL.Query().Get("token_ws"),
		"vci": webpay.Vci,
		"amount": amount,
		"status": webpay.Status,
		"buy_order": webpay.Buy_order,
		"session_id": webpay.Session_id,
		"card_detail": webpay.Card_detail["card_number"],
		"accounting_date": webpay.Accounting_date,
		"transaction_date": webpay.Transaction_date,
		"authorization_code": webpay.Authorization_code,
		"payment_type_code": webpay.Payment_type_code,
		"response_code": response_code,
		"installments_number": installments_number,
	}
	template.Execute(response, data)
}

func Pasarelas_paypal(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/pasarelas/paypal.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)

	// La nueva función se encarga de todo el proceso de PayPal
	paypal_id, token := crearOrdenPaypal()

	data := map[string]string{
		"css": css_sesion,
		"mensaje": css_mensaje,
		"token": token,
		"paypal_id": paypal_id,
	}
	template.Execute(response, data)
}

// crearOrdenPaypal utiliza la librería para un flujo más seguro y robusto.
// Devuelve el ID de la orden y el token de acceso.
func crearOrdenPaypal() (string, string) {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		fmt.Println("Error loading .env file")
		return "", ""
	}

	client, err := paypal.NewClient(os.Getenv("PAYPAL_CLIENT_ID"), os.Getenv("PAYPAL_CLIENT_SECRET"), paypal.APIBaseSandBox)
	if err != nil {
		fmt.Println("Error creating paypal client:", err)
		return "", ""
	}
	client.SetLog(os.Stdout) // Mantenemos el log para depuración

	accessToken, err := client.GetAccessToken(context.Background())
	if err != nil {
		fmt.Println("Error getting access token:", err)
		return "", ""
	}

	// Definimos la orden usando las estructuras de la librería
	purchaseUnits := []paypal.PurchaseUnitRequest{
		{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: "USD",
				Value: "10.00",
			},
			ReferenceID: "orden_1",
		},
	}

	// Usamos ApplicationContext para definir la experiencia de pago.
	appContext := &paypal.ApplicationContext{
		BrandName: "MR_Hola",
		Locale: "es-ES",
		LandingPage: "LOGIN",
		ShippingPreference: "NO_SHIPPING",
		UserAction: "PAY_NOW",
		ReturnURL: "http://localhost:8080/pasarelas/paypal/retorno",
		CancelURL: "http://localhost:8080/pasarelas/paypal/cancelado",
	}

	// Usamos el método de la librería para crear la orden
	order, err := client.CreateOrder(context.Background(), paypal.OrderIntentCapture, purchaseUnits, nil, appContext)
	if err != nil {
		fmt.Println("Error creating order:", err)
		// Imprimir detalles del error de PayPal si están disponibles
		if payPalErr, ok := err.(*paypal.ErrorResponse); ok {
			errorDetails, _ := json.MarshalIndent(payPalErr.Details, "", "  ")
			fmt.Println("PayPal Error Details:", string(errorDetails))
		}
		return "", ""
	}

	return order.ID, accessToken.Token
}

func Pasarelas_paypal_retorno(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/pasarelas/paypal_retorno.html", utilidades.Frontend))

	// Se necesita un cliente de PayPal para trabajar con la API
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		fmt.Println("Error loading .env file:", errorVariables)
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	client, err := paypal.NewClient(os.Getenv("PAYPAL_CLIENT_ID"), os.Getenv("PAYPAL_CLIENT_SECRET"), paypal.APIBaseSandBox)
	if err != nil {
		fmt.Println("Error creating paypal client:", err)
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Capturamos la orden usando el ID (token) que nos devuelve PayPal en la URL
	orderID := request.URL.Query().Get("token")
	captureResponse, err := client.CaptureOrder(context.Background(), orderID, paypal.CaptureOrderRequest{})

	var estado_paypal string
	if err != nil {
		fmt.Println("Error capturing order:", err)
		// Imprimir detalles del error de PayPal si están disponibles
		if payPalErr, ok := err.(*paypal.ErrorResponse); ok {
			errorDetails, _ := json.MarshalIndent(payPalErr.Details, "", "  ")
			fmt.Println("PayPal Capture Error Details:", string(errorDetails))
		}
		estado_paypal = "mal"
	} else {
		// Si la captura es exitosa, el estado debe ser "COMPLETED"
		if captureResponse.Status == "COMPLETED" {
			estado_paypal = "bien"
		} else {
			// Otro estado, ej. el pago requiere verificación adicional
			estado_paypal = "pendiente"
		}
	}

	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css":           css_sesion,
		"mensaje":       css_mensaje,
		"token":         orderID,
		"estado_paypal": estado_paypal,
	}
	template.Execute(response, data)
}

func Pasarelas_paypal_cancelar(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/pasarelas/paypal.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)

	// La nueva función se encarga de todo el proceso de PayPal
	paypal_id, token := crearOrdenPaypal()

	data := map[string]string{
		"css": css_sesion,
		"mensaje": css_mensaje,
		"token": token,
		"paypal_id": paypal_id,
	}
	template.Execute(response, data)
}