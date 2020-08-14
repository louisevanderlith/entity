package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/entity/core"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

var Manager kong.Manager

func SetupRoutes(scrt, securityUrl string) http.Handler {
	Manager = kong.NewManager(core.Context())

	r := mux.NewRouter()

	r.HandleFunc("/login", kong.ResourceMiddleware(http.DefaultClient, "entity.login.apply", scrt, securityUrl, "", LoginPOST)).Methods(http.MethodPost)
	r.HandleFunc("/consent", kong.ResourceMiddleware(http.DefaultClient, "entity.consent.apply", scrt, securityUrl, "", ConsentPOST)).Methods(http.MethodPost)

	get := kong.ResourceMiddleware(http.DefaultClient, "entity.info.search", scrt, securityUrl, "", GetInfo)
	r.HandleFunc("/", get).Methods(http.MethodGet)

	view := kong.ResourceMiddleware(http.DefaultClient, "entity.info.view", scrt, securityUrl, "", ViewInfo)
	r.HandleFunc("/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := kong.ResourceMiddleware(http.DefaultClient, "entity.info.search", scrt, securityUrl, "", SearchInfo)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := kong.ResourceMiddleware(http.DefaultClient, "entity.info.search", scrt, securityUrl, "", CreateInfo)
	r.HandleFunc("/", create).Methods(http.MethodPost)

	r.HandleFunc("/register", kong.ResourceMiddleware(http.DefaultClient, "entity.info.register", scrt, securityUrl, "", RegisterPOST)).Methods(http.MethodPost)

	lst, err := kong.Whitelist(http.DefaultClient, securityUrl, "entity.info.view", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
