package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"log"
	"net/http"

	"github.com/louisevanderlith/entity/core"
)

func GetInsight(w http.ResponseWriter, r *http.Request) {
	idn := drx.GetIdentity(r)

	if !idn.HasUser() {
		log.Println("No User for Insight")
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ut, err := Manager.Insight(idn.GetUserToken())

	if err != nil {
		log.Println("Insight Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(ut))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	results, err := core.Context().GetEntities(1, 10)

	if err != nil {
		log.Println("Get Entities Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title GetEntities
// @Description Gets the entities
// @Success 200 {[]core.Entity} []core.Entity
// @router /all/:pagesize [get]
func SearchInfo(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)

	results, err := core.Context().GetEntities(page, size)

	if err != nil {
		log.Println("Get Entities Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title GetEntity
// @Description Gets the requested Entity
// @Param	key			path	hsk.Key 	true		"Key of the entity you require"
// @Success 200 {core.Entity} core.Entity
// @router /:key [get]
func ViewInfo(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Key Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	record, err := core.Context().GetEntity(key)

	if err != nil {
		log.Println("Get Entity Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(record.GetValue()))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title CreateEntity
// @Description Creates a comment
// @Param	body		body 	logic.Entity	true		"Entity entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateInfo(w http.ResponseWriter, r *http.Request) {
	var entry core.Entity
	err := drx.JSONBody(r, &entry)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ctx := core.Context()
	_, err = ctx.CreateEntity(entry)

	if err != nil {
		log.Println("Create Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON("CREATED"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
