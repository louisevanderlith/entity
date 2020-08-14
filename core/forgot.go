package core

import (
	"github.com/louisevanderlith/husk"
)

//Forgot is used to keep a list of Users which have requested to Change their passwords.
type Forgot struct {
	UserKey  husk.Key
	Redeemed bool
}

func (v Forgot) Valid() error {
	return husk.ValidateStruct(v)
}
