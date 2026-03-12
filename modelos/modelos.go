package modelos

type Cliente struct {
	Id int
	Nombre string
	Correo string
	Telefono string
}

type Clientes []Cliente

type Categoria struct {
	Id int
	Nombre string
	Slug string
}

type Categorias []Categoria

type ClienteHttp struct {
	Css		string
	Mensaje	string
	Datos	Clientes
}
type ClienteHttp2 struct {
	Css		string
	Mensaje	string
	Datos	Cliente
}

type Usuario struct {
	Id int
	Nombre string
	Correo string
	Telefono string
	Password string
}

type Usuarios []Usuario

type WebpayModel struct{
	Url string
	Token string
}

type WebpayRetornoModel struct{
	Vci string
	Amount int
	Status string
	Buy_order string
	Session_id string
	Card_detail map[string]string
	Accounting_date string
	Transaction_date string
	Authorization_code string
	Payment_type_code string
	Response_code int
	Installments_number int
}

type PaypalOrderResponseModel struct{
	Id string
	Status string
	Links string
}