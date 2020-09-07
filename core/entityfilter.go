package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type entityFilter func(obj Entity) bool

func (f entityFilter) Filter(obj hsk.Record) bool {
	return f(obj.Data().(Entity))
}

//Email filter will filter by email and verification status
func byEmail(email string) entityFilter {
	return func(obj Entity) bool {
		return obj.User.Email == email && obj.User.Verified
	}
}
