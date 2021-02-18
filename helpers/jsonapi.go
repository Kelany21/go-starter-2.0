package helpers

type JsonApiData struct {
	Type          string                 `json:"type"`
	Id            string                 `json:"id"`
	Attributes    map[string]interface{} `json:"attributes"`
	Relationships map[string]interface{} `json:"relationships"`
}

type JsonApiIncluded struct {
	Type          string                 `json:"type"`
	Id            string                 `json:"id"`
	Attributes    map[string]interface{} `json:"attributes"`
	Relationships map[string]interface{} `json:"relationships"`
}

type JsonApiMeta struct {
	Filters    map[string][]map[string]interface{} `json:"filters"`
	Pagination map[string]int                      `json:"pagination"`
}

type JsonApi struct {
	Data     []JsonApiData     `json:"data"`
	Included []JsonApiIncluded `json:"included"`
	Meta     JsonApiMeta       `json:"meta"`
	Links    map[string]string `json:"links"`
}
