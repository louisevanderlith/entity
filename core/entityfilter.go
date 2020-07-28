package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/kong/prime"
)

//Email filter will filter by email and verification status
func emailFilter(email string) userFilter {
	return func(obj prime.User) bool {
		return obj.Email == email && obj.Verified
	}
}

