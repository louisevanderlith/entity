package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/entity/core"
	"github.com/louisevanderlith/husk"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	k := husk.CrazyKey()
	result, err := core.Context().GetEntity(k)

	if err != nil {
		log.Println("Get Entity Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	k := husk.CrazyKey()
	result, err := core.Context().GetEntity(k)

	if err != nil {
		log.Println("Get Entity Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
