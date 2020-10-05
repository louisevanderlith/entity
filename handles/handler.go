package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/entity/core"
	"github.com/louisevanderlith/kong"
	"github.com/louisevanderlith/kong/middle"
	"github.com/rs/cors"
	"net/http"
)

var Manager kong.Manager

func SetupRoutes(scrt, securityUrl string) http.Handler {
	Manager = kong.NewManager(core.Context())

	r := mux.NewRouter()
	ins := middle.NewResourceInspector(http.DefaultClient, securityUrl, "")
	r.HandleFunc("/login", ins.Middleware("entity.login.apply", scrt, LoginPOST)).Methods(http.MethodPost)
	r.HandleFunc("/consent", ins.Middleware("entity.consent.apply", scrt, ConsentPOST)).Methods(http.MethodPost)
	r.HandleFunc("/insight", ins.Middleware("entity.info.view", scrt, GetInsight)).Methods(http.MethodGet)

	get := ins.Middleware("entity.info.search", scrt, GetInfo)
	r.HandleFunc("/", get).Methods(http.MethodGet)

	view := ins.Middleware("entity.info.view", scrt, ViewInfo)
	r.HandleFunc("/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := ins.Middleware("entity.info.search", scrt, SearchInfo)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := ins.Middleware("entity.info.search", scrt, CreateInfo)
	r.HandleFunc("/", create).Methods(http.MethodPost)

	r.HandleFunc("/register", ins.Middleware("entity.info.register", scrt, RegisterPOST)).Methods(http.MethodPost)

	lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "entity.info.view", scrt)

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
