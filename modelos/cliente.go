package modelos

type Client struct {
	ID       int64  `json:"id"`
	Nombre   string `json:"nombre"`
	Correo   string `json:"email"`
	Telefono string `json:"phone"`
	Password string `json:"password"`
}

type Clients []Client
