package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"net/http"
)

func ConsentPOST(w http.ResponseWriter, r *http.Request) {
	claims := make(map[string]bool)
	obj := prime.QueryRequest{Claims: &claims}
	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ut, err := Manager.Consent(obj.Token, claims)

	if err != nil {
		log.Println("Consent Error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc, err := Manager.Sign(ut, 10)

	if err != nil {
		log.Println("Sign Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(enc))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
