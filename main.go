package main

import (
	"flag"
	"github.com/louisevanderlith/entity/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/entity/core"
)

func main() {
	securityUrl := flag.String("security", "http://localhost:8086", "Security Provider's URL")
	srcSecrt := flag.String("scopekey", "secret", "Secret used to validate against scopes")

	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8097",
		Handler:      handles.SetupRoutes(*srcSecrt, *securityUrl),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
