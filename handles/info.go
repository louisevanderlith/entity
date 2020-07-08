package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/entity/core"
	"github.com/louisevanderlith/husk"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	results, err := core.GetEntities(1, 10)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println(err)
	}
}

// @Title GetEntities
// @Description Gets the entities
// @Success 200 {[]core.Entity} []core.Entity
// @router /all/:pagesize [get]
func SearchInfo(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	page, size := ctx.GetPageData()

	results, err := core.GetEntities(page, size)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println(err)
	}
}

// @Title GetEntity
// @Description Gets the requested Entity
// @Param	key			path	husk.Key 	true		"Key of the entity you require"
// @Success 200 {core.Entity} core.Entity
// @router /:key [get]
func ViewInfo(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	record, err := core.GetEntity(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(record.Data()))

	if err != nil {
		log.Println(err)
	}
}

// @Title CreateEntity
// @Description Creates a comment
// @Param	body		body 	logic.Entity	true		"Entity entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateInfo(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	var entry core.Entity
	err := ctx.Body(&entry)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := entry.Create()

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec.Data()))

	if err != nil {
		log.Println(err)
	}
}
