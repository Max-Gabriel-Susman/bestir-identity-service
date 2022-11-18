package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/account"
	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/foundation/web"
)

var _ http.Handler = (*web.App)(nil)

// maybe we'll add gitsha and other params later
func API() *web.App {
	app := web.NewApp()

	ag := account.AccountGroup{}
	app.Handle("GET", "/accounts", ag.GetAccounts)

	// we'll want more robust registry down hither at a later point in time

	return app
}

// this is just a temporary measure so we can get some rudimentary shit up and running
func Handler() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
	return nil
}
