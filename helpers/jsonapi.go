package helpers

import (
	"github.com/google/uuid"
	"strconv"
	"strings"
)

type Relationship struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Relationships struct {
	Data []Relationship `json:"data"`
}

type JsonApiData struct {
	Type          string                   `json:"type"`
	Id            string                   `json:"id"`
	Attributes    map[string]interface{}   `json:"attributes"`
	Relationships map[string]Relationships `json:"relationships"`
}

type JsonApiIncluded struct {
	Type          string                   `json:"type"`
	Id            string                   `json:"id"`
	Attributes    map[string]interface{}   `json:"attributes"`
	Relationships map[string]Relationships `json:"relationships"`
}

type JsonApiFilter struct {
	FilterKey string                 `json:"filter_key"`
	Value     map[string]interface{} `json:"value"`
	ValueKeys []string               `json:"value_keys"`
}

type JsonApiMeta struct {
	Filters    map[string][]JsonApiFilter `json:"filters"`
	Pagination map[string]int             `json:"pagination"`
}

type JsonApi struct {
	Data     []JsonApiData     `json:"data"`
	Included []JsonApiIncluded `json:"included"`
	Meta     JsonApiMeta       `json:"meta"`
	Links    map[string]string `json:"links"`
}

type JsonApiPrepare struct {
	Includes   map[string]interface{}
	UUID       string
	Attributes map[string]interface{}
}

func PrepareResponse(jsonapiData *[]JsonApiData, jsonapiIncluded *[]JsonApiIncluded, prepareObject JsonApiPrepare) {
	relations := make(map[string]Relationships)
	for includeType, includes := range prepareObject.Includes {
		var relation Relationships
		for _, value := range includes.([]map[string]interface{}) {
			attributes := make(map[string]interface{})
			if value["attributes"] != nil {
				attributes = value["attributes"].(map[string]interface{})
			} else {
				attributes = nil
			}
			relationships := make(map[string]Relationships)
			if value["relationships"] != nil {
				relationships = value["relationships"].(map[string]Relationships)
			} else {
				relationships = nil
			}
			include := JsonApiIncluded{
				Type:          includeType,
				Id:            value["id"].(uuid.UUID).String(),
				Attributes:    attributes,
				Relationships: relationships,
			}
			*jsonapiIncluded = append(*jsonapiIncluded, include)

			relation.Data = append(relation.Data, Relationship{
				Type: include.Type,
				Id:   include.Id,
			})
		}
		relations[includeType] = relation
	}
	*jsonapiData = append(*jsonapiData, JsonApiData{
		Type:          "user",
		Id:            prepareObject.UUID,
		Attributes:    prepareObject.Attributes,
		Relationships: relations,
	})
}

func PaginationObject(paginator *Paginator) map[string]int {
	var u = make(map[string]int)
	u["total"] = paginator.TotalRecord
	u["count"] = paginator.RecordsCount
	u["per_page"] = paginator.Limit
	u["current_page"] = paginator.Page
	u["total_pages"] = paginator.TotalPage

	return u
}

func PaginationLinks(paginator *Paginator, url string) map[string]string {
	if strings.Contains(url, "?") {
		url += "&page="
	} else {
		url += "?page="
	}
	limitQuery := "&limit=" + strconv.Itoa(paginator.Limit)
	var u = make(map[string]string)
	u["self"] = url + strconv.Itoa(paginator.Page) + limitQuery
	u["next"] = url + strconv.Itoa(paginator.NextPage) + limitQuery
	u["prev"] = url + strconv.Itoa(paginator.PrevPage) + limitQuery
	u["first"] = url + "1" + limitQuery
	u["last"] = url + strconv.Itoa(paginator.TotalPage) + limitQuery

	return u
}
