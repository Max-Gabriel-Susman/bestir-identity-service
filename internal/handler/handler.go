package handler

import (
	"net/http"

	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/account"
	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/foundation/web"
)

var _ http.Handler = (*web.App)(nil)

// maybe we'll add gitsha and other params later
func API() *web.App {
	app := web.NewApp()

	accountAPI := account.NewAPI()
	// userAPI := user.NewAPI()

	// we'll want more robust registry down hither at a later point in time

	// UserEndpoints(app, userAPI)
	AccountEndpoints(app, accountAPI)
	return app
}
