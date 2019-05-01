package core

import "github.com/louisevanderlith/husk"

type context struct {
	Entities husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Entities: husk.NewTable(new(Entity)),
	}
}

func Shutdown() {
	ctx.Entities.Save()
}
