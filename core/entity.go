package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/kong/prime"
)

type Entity struct {
	Name           string `hsk:"size(30)"`
	ProfileKey     string
	User           prime.User
	Identification string    `hsk:"size(30)"` //This can be a Company Registration, RSA ID, Passport or another interal identifier
	Addresses      []Address `hsk:"null"`
}

func (e Entity) Valid() error {
	return husk.ValidateStruct(e)
}

func (e Entity) Create() (husk.Recorder, error) {
	rec, err := ctx.Entities.Create(e)

	if err != nil {
		return nil, err
	}

	err = ctx.Entities.Save()

	if err != nil {
		return nil, err
	}

	return rec, nil
}
