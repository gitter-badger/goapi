package resources

import "github.com/exotel/goapi/assets/types"

//AccountDetails struct for AccountDetails request data
type AccountDetails struct {
	FriendlyName string `mandatory:"true" url:"FriendlyName"`
	CountryCode  string `mandatory:"true" queryparam:"CountryCode"`
}

//AccountFilter struct for AccountFilter request data
type AccountFilter struct {
	FriendlyName string
	Status       string
}

//AccountUpdatableDetails struct for Update request data
type AccountUpdatableDetails struct {
	FriendlyName string
	Status       string
}

//AccountURLS saves the routes its in the format that the text/template library of
//go can accept
var AccountURLS = map[types.Action]string{
	types.CREATE: "/v1/Accounts/{{.FriendlyName}}",
	types.READ:   "/v1/Accounts/{{.ResourceID}}",
	types.UPDATE: "/v1/Accounts/{{.ResourceID}}",
	types.DELETE: "/v1/Accounts/{{.ResourceID}}",
}