package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxo"
	"github.com/louisevanderlith/squareroot/ctx"
	"net/http"

	"github.com/louisevanderlith/entity/core"
	"github.com/louisevanderlith/husk"
)

func Get(c *gin.Context) {
	results := core.GetEntities(1, 10)

	c.JSON(http.StatusOK, results)
}

// @Title GetEntities
// @Description Gets the entities
// @Success 200 {[]core.Entity} []core.Entity
// @router /all/:pagesize [get]
func Search(c *gin.Context) {
	page, size := droxo.GetPageData(c.Param("pagesize"))

	results := core.GetEntities(page, size)

	c.JSON(http.StatusOK, results)
}

// @Title GetEntity
// @Description Gets the requested Entity
// @Param	key			path	husk.Key 	true		"Key of the entity you require"
// @Success 200 {core.Entity} core.Entity
// @router /:key [get]
func View(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	record, err := core.GetEntity(key)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, record)
}

// @Title CreateEntity
// @Description Creates a comment
// @Param	body		body 	logic.Entity	true		"Entity entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func Create(c *gin.Context) {
	var entry core.Entity
	err := ctx.Body(&entry)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rec, err := entry.Create()

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, rec)
}
