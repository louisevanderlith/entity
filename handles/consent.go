package handles

import (
	"encoding/json"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"net/http"
)

func ConsentPOST(w http.ResponseWriter, r *http.Request) {
	obj := prime.QueryRequest{}
	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ut, err := Manager.Consent(obj.Token, obj.Claims.(map[string]bool))

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

	bits, err := json.Marshal(enc)

	if err != nil {
		log.Println("Marshal Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(bits)

	if err != nil {
		log.Println("Serve Error", err)
	}
}
