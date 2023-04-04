package accessmanagerclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Result                  []Realm     `json:"result"`
	ResultCount             int         `json:"resultCount"`
	PagedResultsCookie      interface{} `json:"pagedResultsCookie"`
	TotalPagedResultsPolicy string      `json:"totalPagedResultsPolicy"`
	TotalPagedResults       int         `json:"totalPagedResults"`
	RemainingPagedResults   int         `json:"remainingPagedResults"`
}

// GetRealms - Returns a list of realms (auth required)
func (c *Client) GetRealms() (*Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/am/json/global-config/realms?_queryFilter=true", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	r, err := c.DoRequest(req)
	if err != nil {
		return nil, err
	}

	response := Response{}

	if err := json.Unmarshal(r, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
