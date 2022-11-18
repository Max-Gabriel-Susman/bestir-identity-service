package account

type API struct {
	// Logger *bestirlog.Logger
	// Store *CockroachDBStorage
}

// we may want to parameterize storage and logging later
func NewAPI() *API {
	return &API{}
}
