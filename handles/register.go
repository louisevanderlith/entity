package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/entity/core"
	"log"
	"net/http"
)

// @Title Register
// @Description Registers a new user
// @Param	body		body 	core.AuthRequest		true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	var regis core.Registration
	err := drx.JSONBody(r, &regis)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.Context().Register(regis)

	if err != nil {
		log.Println("Register Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON("SAVED"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
