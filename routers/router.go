package routers

// @APIVersion 1.0.0
// @Title Entity.API
// @Description API used to access and modify enity details.

import (
	"github.com/louisevanderlith/entity/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
)

func Setup(poxy resins.Epoxi) {
	//Info
	infoCtrl := &controllers.InfoController{}
	infoGroup := routing.NewRouteGroup("info", mix.JSON)
	infoGroup.AddRoute("Create Information", "", "POST", roletype.Owner, infoCtrl.Post)
	infoGroup.AddRoute("All Information", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, infoCtrl.Get)
	infoGroup.AddRoute("View Information", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, infoCtrl.GetByID)
	poxy.AddGroup(infoGroup)
	/*ctrlmap := EnableFilters(s, host)
	infoctrl := controllers.NewInfoCtrl(ctrlmap)

	beego.Router("/v1/info", infoctrl, "post:Post")
	beego.Router("/v1/info/:key", infoctrl, "get:GetByID")
	beego.Router("/v1/info/all/:pagesize", infoctrl, "get:Get")*/
}

/*
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
*/
