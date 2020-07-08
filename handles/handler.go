package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	r := mux.NewRouter()

	get := kong.ResourceMiddleware("entity.info.search", scrt, secureUrl, GetInfo)
	r.HandleFunc("/", get).Methods(http.MethodGet)

	view := kong.ResourceMiddleware("entity.info.view", scrt, secureUrl, ViewInfo)
	r.HandleFunc("/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := kong.ResourceMiddleware("entity.info.search", scrt, secureUrl, SearchInfo)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := kong.ResourceMiddleware("blog.articles.create", scrt, secureUrl, CreateInfo)
	r.HandleFunc("/", create).Methods(http.MethodPost)

	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "entity.info.view", scrt)

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
