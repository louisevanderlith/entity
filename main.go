package main

import (
	"log"
	"os"

	"github.com/louisevanderlith/entity/core"
	"github.com/louisevanderlith/entity/routers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"

	"github.com/astaxie/beego"
)

func main() {
	mode := os.Getenv("RUNMODE")
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	pubPath := path.Join(keyPath, pubName)

	core.CreateContext()
	defer core.Shutdown()

	// Register with router
	name := beego.BConfig.AppName
	srv := mango.NewService(mode, name, pubPath, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
