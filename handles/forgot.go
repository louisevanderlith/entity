package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/entity/core"
)

// @Title Forgot Password
// @Description Will send the user an email with an OTP
// @Param	body		body 	logic.Login	true		"body for message content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func CreateForgot(w http.ResponseWriter, r *http.Request) {
	email := ""
	err := drx.JSONBody(r, &email)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	resp, err := core.Context().RequestReset(email, r.URL.RequestURI())

	if err != nil {
		log.Println("Request Reset Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(resp))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
