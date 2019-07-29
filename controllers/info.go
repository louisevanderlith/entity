package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/entity/core"
	"github.com/louisevanderlith/husk"
)

type InfoController struct {
	xontrols.APICtrl
}

// @Title GetEntities
// @Description Gets the entities
// @Success 200 {[]core.Entity} []core.Entity
// @router /all/:pagesize [get]
func (req *InfoController) Get() {
	page, size := req.GetPageData()

	results := core.GetEntities(page, size)

	req.Serve(http.StatusOK, nil, results)
}

// @Title GetEntity
// @Description Gets the requested Entity
// @Param	key			path	husk.Key 	true		"Key of the entity you require"
// @Success 200 {core.Entity} core.Entity
// @router /:key [get]
func (req *InfoController) GetByID() {
	key, err := husk.ParseKey(req.FindParam("key"))

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	record, err := core.GetEntity(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, record)
}

// @Title CreateEntity
// @Description Creates a comment
// @Param	body		body 	logic.Entity	true		"Entity entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *InfoController) Post() {
	var entry core.Entity
	err := req.Body(&entry)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec, err := entry.Create()

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}
