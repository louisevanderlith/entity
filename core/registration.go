package core

import (
	"errors"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/kong/prime"
)

type Registration struct {
	Name           string
	Email          string
	Password       string
	PasswordRepeat string
	ProfileClient  string
}

func Register(r Registration) (hsk.Key, error) {
	if r.Password != r.PasswordRepeat {
		return nil, errors.New("passwords do not match")
	}

	if ctx.Entities.Exists(byEmail(r.Email)) {
		return nil, errors.New("email already in use")
	}

	contc := prime.Contacts{
		{
			Icon:  "fa-mail",
			Name:  "email",
			Value: r.Email,
		},
	}

	//TODO: Make dynamic
	//Should provide only basic Resources, the rest will be unlocked later
	user := prime.NewUser(r.Name, r.Email, r.Password, false, contc, nil)

	return ctx.Entities.Create(user.(prime.User))
}
