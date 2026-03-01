package rutas

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"proyecto_3/modelos"
	"proyecto_3/utilidades"
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
	qrcode "github.com/skip2/go-qrcode"
	excelize "github.com/xuri/excelize/v2"
)

func Recursos_utiles_get(response http.ResponseWriter, request *http.Request) { 
	template := template.Must(template.ParseFiles("templates/recursos_utiles/home.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css" :		css_sesion,
		"mensaje" : css_mensaje,
	}
	template.Execute(response, data)
}

// ------------- PDF -------------

func ImageFile(fileStr string) string {
	return filepath.Join(gofpdfDir, "static/images", fileStr)
}

var gofpdfDir string

func Filename(baseStr string) string {
	return PdfFile(baseStr + ".pdf")
}
func PdfFile(fileStr string) string {
	return filepath.Join(PdfDir(), fileStr)
}
func PdfDir() string {
	return filepath.Join(gofpdfDir, "static/pdf")
}
func Summary(err error, fileStr string) {
	if err == nil {
		fileStr = filepath.ToSlash(fileStr)
		fmt.Printf("Successfully generated %s\n", fileStr)
	} else {
		fmt.Println(err)
	}
}

func Recursos_utiles_pdf(response http.ResponseWriter, request *http.Request) { 
	template := template.Must(template.ParseFiles("templates/recursos_utiles/pdf.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css" :		css_sesion,
		"mensaje" : css_mensaje,
	}
	template.Execute(response, data)
}

func Recursos_utiles_pdf_generar(response http.ResponseWriter, request *http.Request) { 
	pdf := gofpdf.New("P", "mm", "A4", "")
	// First page: manual local link
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 20)
	_, lineHt := pdf.GetFontSize()
	pdf.Write(lineHt, "To find out what's new in this tutorial, click ")
	pdf.SetFont("", "U", 0)
	link := pdf.AddLink()
	pdf.WriteLinkID(lineHt, "here", link)
	pdf.SetFont("", "", 0)
	// Second page: image link and basic HTML with link
	pdf.AddPage()
	pdf.SetLink(link, 0, -1)
	pdf.Image(ImageFile("logo.png"), 10, 12, 30, 0, false, "", 0, "http://www.fpdf.org")
	pdf.SetLeftMargin(45)
	pdf.SetFontSize(14)
	_, lineHt = pdf.GetFontSize()
	htmlStr := `You can now easily print text mixing different styles: <b>bold</b>, ` +
		`<i>italic</i>, <u>underlined</u>, or <b><i><u>all at once</u></i></b>!<br><br>` +
		`<center>You can also center text.</center>` +
		`<right>Or align it to the right.</right>` +
		`You can also insert links on text, such as ` +
		`<a href="http://www.fpdf.org">www.fpdf.org</a>, or on an image: click on the logo.`
	html := pdf.HTMLBasicNew()
	html.Write(lineHt, htmlStr)
	time := strings.Split(time.Now().String(), " ")
	nombre := string(time[4][6:14])
	fileStr := Filename(nombre)
	err := pdf.OutputFileAndClose(fileStr)
	Summary(err, fileStr)

	mensaje := "Se creó el documento PDF " + nombre + ".pdf de forma correcta"

	utilidades.CrearMensajesFlash(response, request, "success", mensaje)
	http.Redirect(response, request, "/recursos-utiles/pdf", http.StatusSeeOther)
}

func Recursos_utiles_excel(response http.ResponseWriter, request *http.Request) { 
	template := template.Must(template.ParseFiles("templates/recursos_utiles/excel.html", utilidades.Frontend))
	//excel
	f := excelize.NewFile()
	defer func ()  {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.SetCellValue("Sheet1", "A1", "id")
	f.SetCellValue("Sheet1", "B1", "nombre")
	f.SetCellValue("Sheet1", "C1", "correo")
	f.SetActiveSheet(index)
	// Se agregan datos
	cliente1 := modelos.Clientes{
		modelos.Cliente{Id: 1, Nombre: "Júan", Correo: "juan@ejemplo.com", Telefono: ""}, 
		modelos.Cliente{Id: 2, Nombre: "Maria ñ", Correo: "maria@ejemplo.com", Telefono: ""},
	}

	contador := 2 // Contador para filas, inicia en 2 porque la 1 es para los titulos, es un int
	i := 0
	for _, service := range cliente1 {
		fila := strconv.Itoa(contador)
		f.SetCellValue("Sheet1", "A" + fila, service.Id)
		f.SetCellValue("Sheet1", "B" + fila, service.Nombre)
		f.SetCellValue("Sheet1", "C" + fila, service.Correo)
		contador++
		i++
	}

	// Se contruye el archivo excel
	time := strings.Split(time.Now().String(), " ")
	nombre := string(time[4][6:14]) + ".xlsx"
	if err := f.SaveAs("static/excel/" + nombre); err != nil {
		fmt.Println(err)
	}

	// Retornar
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css" :		css_sesion,
		"mensaje" : css_mensaje,
		"nombre" : nombre,
	}
	template.Execute(response, data)
}


func Recursos_utiles_qr(response http.ResponseWriter, request *http.Request) { 
	template := template.Must(template.ParseFiles("templates/recursos_utiles/qr.html", utilidades.Frontend))
	// Generar QR
	ImagenCodificada, err := qrcode.Encode("https://github.com/brzrker", qrcode.High, 256)
	if err != nil {
		log.Fatalln("Error al generar el código qr", err)
	}
	imagenQR := base64.StdEncoding.EncodeToString(ImagenCodificada)

	// Retornar
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css" :		css_sesion,
		"mensaje" : css_mensaje,
		"imagenQR" : imagenQR,
	}
	template.Execute(response, data)
}

func Recursos_utiles_enviar_correo(response http.ResponseWriter, request *http.Request) { 
	template := template.Must(template.ParseFiles("templates/recursos_utiles/enviar_correo.html", utilidades.Frontend))
	utilidades.EnviarCorreo()
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{
		"css" :		css_sesion,
		"mensaje" : css_mensaje,
	}
	template.Execute(response, data)
}