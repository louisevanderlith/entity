package core

import (
	"github.com/louisevanderlith/husk"
	"log"
)

type EntityContext interface {
	GetEntities(page, pageSize int) (husk.Collection, error)
	GetEntity(k husk.Key) (husk.Recorder, error)
}

type context struct {
	Entities husk.Tabler
}

var ctx context

func Context() EntityContext {
	return ctx
}

func CreateContext() {
	ctx = context{
		Entities: husk.NewTable(Entity{}),
	}
}

func Shutdown() {
	ctx.Entities.Save()
}

func (c context) GetEntities(page, pageSize int) (husk.Collection, error) {
	return c.Entities.Find(page, pageSize, husk.Everything())
}

func (c context) GetEntity(k husk.Key) (husk.Recorder, error) {
	return c.Entities.FindByKey(k)
}

func (c context) GetUser(page, size int) []SafeUser {
	var result []SafeUser
	users, err := c.Entities.Find(page, size, husk.Everything())

	if err != nil {
		log.Println(err)
		return nil
	}

	itor := users.GetEnumerator()

	for itor.MoveNext() {
		currUser := itor.Current()

		sfeUser := createSafeUser(currUser)
		result = append(result, sfeUser)
	}

	return result
}