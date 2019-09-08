package routers

// @APIVersion 1.0.0
// @Title Entity.API
// @Description API used to access and modify enity details.

import (
	"github.com/louisevanderlith/entity/controllers"

	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	e.JoinBundle("/", roletype.Owner, mix.JSON, &controllers.Info{})
	//Info
	/*infoCtrl := &controllers.InfoController{}
	infoGroup := routing.NewRouteGroup("info", mix.JSON)
	infoGroup.AddRoute("Create Information", "", "POST", roletype.Owner, infoCtrl.Post)
	infoGroup.AddRoute("All Information", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.User, infoCtrl.Get)
	infoGroup.AddRoute("View Information", "/{key:[0-9]+\x60[0-9]+}", "GET", roletype.User, infoCtrl.GetByID)
	e.AddBundle(infoGroup)*/
}
