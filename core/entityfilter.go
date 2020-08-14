package core

import (
	"github.com/louisevanderlith/husk"
)

type entityFilter func(obj Entity) bool

func (f entityFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(Entity))
}

//Email filter will filter by email and verification status
func byEmail(email string) entityFilter {
	return func(obj Entity) bool {
		return obj.User.Email == email && obj.User.Verified
	}
}
