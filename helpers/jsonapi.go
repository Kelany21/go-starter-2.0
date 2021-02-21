package helpers

import (
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
	var u = make(map[string]string)
	u["self"] = url + strconv.Itoa(paginator.Page)
	u["next"] = url + strconv.Itoa(paginator.NextPage)
	u["prev"] = url + strconv.Itoa(paginator.PrevPage)
	u["first"] = url + "1"
	u["last"] = url + strconv.Itoa(paginator.TotalPage)

	return u
}
