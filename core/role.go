package core

import (
	"github.com/louisevanderlith/husk"
)

type Role struct {
	Application string
	Scope       int
}

func (o Role) Valid() error {
	return husk.ValidateStruct(o)
}
