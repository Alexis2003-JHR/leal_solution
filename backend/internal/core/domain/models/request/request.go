package request

type CreateUser struct {
	Name  string `json:"nombre"`
	Email string `json:"correo"`
}
