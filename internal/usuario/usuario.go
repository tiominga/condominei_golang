package usuario

import (
	"errors"
	"net/mail"
	"strings"
	"time"
)

type Usuario struct {
    ID        uint64
    Email     string
    Senha     string
    Status    int
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (u Usuario) Validar() error {
	if strings.TrimSpace(u.Email) == "" {
		return errors.New("o e-mail é obrigatório")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("o e-mail informado é inválido")
	}

	if strings.TrimSpace(u.Senha) == "" {
		return errors.New("a senha é obrigatória")
	}

	if len(u.Senha) < 8 {
		return errors.New("a senha deve ter pelo menos 8 caracteres")
	}

	if u.Status != 0 && u.Status != 1 {
		return errors.New("status inválido")
	}

	return nil
}

func (u Usuario) BemVindo() string {

	return "Bem vindo " + u.Email

} 