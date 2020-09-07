package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/kong/prime"
	"time"
)

type SafeUser struct {
	Key         hsk.Key
	Name        string
	Verified    bool
	DateCreated time.Time
}

func createSafeUser(k hsk.Key, user prime.User) SafeUser {
	return SafeUser{
		Key:         k,
		Name:        user.Name,
		Verified:    user.Verified,
		DateCreated: k.GetTimestamp(),
	}
}
