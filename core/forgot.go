package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

//Forgot is used to keep a list of Users which have requested to Change their passwords.
type Forgot struct {
	EntityKey  hsk.Key
	Redeemed bool
}

func (v Forgot) Valid() error {
	return validation.Struct(v)
}
