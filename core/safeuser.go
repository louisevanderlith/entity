package core

import (
	"github.com/louisevanderlith/kong/prime"
	"time"

	"github.com/louisevanderlith/husk"
)

type SafeUser struct {
	Key         husk.Key
	Name        string
	Verified    bool
	DateCreated time.Time
}

func createSafeUser(k husk.Key, user prime.User) SafeUser {
	return SafeUser{
		Key:         k,
		Name:        user.Name,
		Verified:    user.Verified,
		DateCreated: time.Unix(0, k.Stamp),
	}
}
