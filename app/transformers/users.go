package transformers

import (
	"github.com/google/uuid"
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/helpers"
)

/**
* stander user response
 */
func LoginResponse(user models.User) map[string]interface{} {
	var u = make(map[string]interface{})
	u["name"] = user.Name
	u["email"] = user.Email
	u["role"] = user.Role
	u["token"] = user.Token
	u["block"] = user.Block

	return u
}

/**
* stander user response
 */
func UserTransform(user models.User) map[string]interface{} {
	var u = make(map[string]interface{})
	u["name"] = user.Name
	u["email"] = user.Email
	u["role"] = user.Role
	u["token"] = user.Token
	u["block"] = user.Block

	return u
}

/**
* user included data
 */
func UserIncludes(user models.User) map[string]interface{} {
	var u = make(map[string]interface{})
	u["action"] = UserActions(user)

	return u
}

/**
* user included actions
 */
func UserActions(user models.User) []map[string]interface{} {
	u := []map[string]interface{}{
		{
			"id": uuid.New(),
			"attributes": map[string]interface{}{
				"endpoint_url": ";jobnvdf" + user.UUID.String(),
				"method":       "GET",
				"label":        "vd",
				"key":          "vcxv",
			},
			"relationships": nil,
		},
	}

	return u
}

/**
* stander the Multi users response
 */
func UsersResponse(users []models.User, paginator *helpers.Paginator, filters map[string][]helpers.JsonApiFilter) helpers.JsonApi {
	var jsonapiData []helpers.JsonApiData
	var jsonapiIncluded []helpers.JsonApiIncluded
	for _, user := range users {
		prepareObject := helpers.JsonApiPrepare{
			Includes:   UserIncludes(user),
			UUID:       user.UUID.String(),
			Attributes: UserTransform(user),
		}
		helpers.PrepareResponse(&jsonapiData, &jsonapiIncluded, prepareObject)
	}
	return helpers.JsonApi{
		Data:     jsonapiData,
		Included: jsonapiIncluded,
		Meta: helpers.JsonApiMeta{
			Filters:    filters,
			Pagination: helpers.PaginationObject(paginator),
		},
		Links: helpers.PaginationLinks(paginator, "http://127.0.0.1:9090/admin/user/paginate"),
	}

}

/**
* stander the Multi users response
 */
func UserResponse(user models.User) helpers.JsonApi {
	var jsonapiData []helpers.JsonApiData
	var jsonapiIncluded []helpers.JsonApiIncluded
	prepareObject := helpers.JsonApiPrepare{
		Includes:   UserIncludes(user),
		UUID:       user.UUID.String(),
		Attributes: UserTransform(user),
	}
	helpers.PrepareResponse(&jsonapiData, &jsonapiIncluded, prepareObject)
	return helpers.JsonApi{
		Data:     jsonapiData,
		Included: jsonapiIncluded,
		Meta: helpers.JsonApiMeta{
			Filters:    nil,
			Pagination: nil,
		},
		Links: nil,
	}

}
