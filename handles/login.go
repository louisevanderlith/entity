package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"net/http"
)

// @Title Login
// @Description Attempts to login against the provided credentials
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func LoginPOST(w http.ResponseWriter, r *http.Request) {
	obj := prime.LoginRequest{}
	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	tkn, err := Manager.Login(obj.Client, obj.Username, obj.Password)

	if err != nil {
		log.Println("Login Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	stkn, err := Manager.Sign(tkn, 5)

	if err != nil {
		log.Println("Sign Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Write([]byte(stkn))
}
