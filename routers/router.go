package routers

// @APIVersion 1.0.0
// @Title Entity.API
// @Description API used to access and modify enity details.

import (
	"fmt"
	"strings"

	"github.com/louisevanderlith/entity/controllers"
	"github.com/louisevanderlith/mango"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service, host string) {
	ctrlmap := EnableFilters(s, host)
	infoctrl := controllers.NewInfoCtrl(ctrlmap)

	beego.Router("/v1/info", infoctrl, "post:Post")
	beego.Router("/v1/info/:key", infoctrl, "get:GetByID")
	beego.Router("/v1/info/all/:pagesize", infoctrl, "get:Get")
}

func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Owner
	emptyMap["GET"] = roletype.User

	ctrlmap.Add("/v1/info", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
