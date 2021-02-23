package transformers

import (
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/helpers"
)

/**
* stander the Multi users response
 */
func UsersResponse(users []models.User, paginator *helpers.Paginator, filters map[string][]helpers.JsonApiFilter) helpers.JsonApi {
	var jsonapiData []helpers.JsonApiData
	var jsonapiIncluded []helpers.JsonApiIncluded
	for _, user := range users {
		helpers.PrepareItemsResponse(&jsonapiData, &jsonapiIncluded, user)
	}
	return helpers.ItemsResponse(jsonapiData, jsonapiIncluded, paginator, filters)
}

/**
* stander the Multi users response
 */
func UserResponse(user models.User) helpers.JsonApi {
	return helpers.ItemResponse(user)
}
