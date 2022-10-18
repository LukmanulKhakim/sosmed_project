package delivery

import (
	"be_project2team4/feature/user/domain"
)

func SuccessResponses(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponses(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

func SuccessLoginResponses(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

type registerRespons struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Bio   string `json:"bio"`
}

type loginResponses struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Bio   string `json:"bio"`
	Token string `json:"token"`
}

type editUserRespons struct {
	Name   string `json:"name"`
	HP     string `json:"hp"`
	Addres string `json:"addres"`
}

func ToResponse(core interface{}, code string, token string) interface{} {
	var res interface{}

	switch code {
	case "reg":
		cnv := core.(domain.Core)
		res = registerRespons{ID: cnv.ID, Name: cnv.Name, Email: cnv.Email, Phone: cnv.Phone, Bio: cnv.Bio}
	case "login":
		cnv := core.(domain.Core)
		res = loginResponses{Name: cnv.Name, Email: cnv.Email, Token: token}
	case "edit":
		cnv := core.(domain.Core)
		res = editUserRespons{Name: cnv.Name}
	}
	return res
}
