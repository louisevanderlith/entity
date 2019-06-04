package core

import "github.com/louisevanderlith/husk"

type Contact struct {
	Email string `hsk:"size(128)" json:",omitempty"`
	Phone string `hsk:"size(20)" json:",omitempty"`
}

func (c Contact) Valid() (bool, error) {
	return husk.ValidateStruct(&c)
}
