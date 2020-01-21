package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Entities husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Entities: husk.NewTable(Entity{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Entities.Save()
}
