package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/entity/core"
	"github.com/louisevanderlith/husk"
)

type Info struct {
}

func (x *Info) Get(ctx context.Requester) (int, interface{}) {
	results := core.GetEntities(1, 10)

	return http.StatusOK, results
}

// @Title GetEntities
// @Description Gets the entities
// @Success 200 {[]core.Entity} []core.Entity
// @router /all/:pagesize [get]
func (req *Info) Search(ctx context.Requester) (int, interface{}) {
	page, size := ctx.GetPageData()

	results := core.GetEntities(page, size)

	return http.StatusOK, results
}

// @Title GetEntity
// @Description Gets the requested Entity
// @Param	key			path	husk.Key 	true		"Key of the entity you require"
// @Success 200 {core.Entity} core.Entity
// @router /:key [get]
func (req *Info) View(ctx context.Requester) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	record, err := core.GetEntity(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, record
}

// @Title CreateEntity
// @Description Creates a comment
// @Param	body		body 	logic.Entity	true		"Entity entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *Info) Create(ctx context.Requester) (int, interface{}) {
	var entry core.Entity
	err := ctx.Body(&entry)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec, err := entry.Create()

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, rec
}
