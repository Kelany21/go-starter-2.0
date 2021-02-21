package transformers

import (
	"github.com/google/uuid"
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/helpers"
)

/**
* stander user response
 */
func UserResponse(user models.User) map[string]interface{} {
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
		relations := make(map[string]helpers.Relationships)
		for includeType, includes := range UserIncludes(user) {
			var relation helpers.Relationships
			for _, value := range includes.([]map[string]interface{}) {
				attributes := make(map[string]interface{})
				if value["attributes"] != nil {
					attributes = value["attributes"].(map[string]interface{})
				}
				relationships := make(map[string]helpers.Relationships)
				if value["relationships"] != nil {
					relationships = value["relationships"].(map[string]helpers.Relationships)
				}
				include := helpers.JsonApiIncluded{
					Type:          includeType,
					Id:            value["id"].(uuid.UUID).String(),
					Attributes:    attributes,
					Relationships: relationships,
				}
				jsonapiIncluded = append(jsonapiIncluded, include)

				relation.Data = append(relation.Data, helpers.Relationship{
					Type: include.Type,
					Id:   include.Id,
				})
			}
			relations[includeType] = relation
		}
		jsonapiData = append(jsonapiData, helpers.JsonApiData{
			Type:          "user",
			Id:            user.UUID.String(),
			Attributes:    UserResponse(user),
			Relationships: relations,
		})
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
