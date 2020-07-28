package handles

import (
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/secure/core"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	result := core.Context.GetUser()

	err := ctx.Serve(http.StatusOK, mix.JSON(result))

	if err != nil {
		log.Println(err)
	}
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	result := core.GetUsers(1, 10)

	err := ctx.Serve(http.StatusOK, mix.JSON(result))

	if err != nil {
		log.Println(err)
	}
}