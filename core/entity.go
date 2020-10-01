package core

import (
	"github.com/louisevanderlith/husk/validation"
	"github.com/louisevanderlith/kong/prime"
)

type Entity struct {
	Name           string `hsk:"size(30)"`
	ProfileKey     string
	User           prime.User
	Identification string    `hsk:"size(30)"` //This can be a Company Registration, RSA ID, Passport or another internal identifier
	Addresses      []Address `hsk:"null"`
}

func (e Entity) Valid() error {
	return validation.Struct(e)
}
