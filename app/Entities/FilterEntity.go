package Entities


type FilterEntity struct {
	MinAmount int `schema:"minAmount"`
	MaxAmount int `schema:"maxAmount"`
	Currency string `schema:"currency"`
	Provider string `schema:"provider"`
	StatusCode string `schema:"statusCode"`
}
