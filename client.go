package accessmanagerclient

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default environment URL
const HostURL string = "https://dev.example.com"

// Client -
type Client struct {
	HostURL         string
	HTTPClient      *http.Client
	amadminSsotoken string
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"xopenamusername"`
	Password string `json:"xopenampassword"`
}

// AuthResponse -
type AuthResponse struct {
	TokenID    string `json:"tokenId"`
	SuccessURL string `json:"successUrl"`
	Realm      string `json:"realm"`
}

// TODO: Move to its own models file
type Realm struct {
	ID         string   `json:"_id"`
	Rev        string   `json:"_rev"`
	ParentPath string   `json:"parentPath"`
	Active     bool     `json:"active"`
	Name       string   `json:"name"`
	Aliases    []string `json:"aliases"`
}

// TODO: Move to its own models file
type Response struct {
	Result                  []Realm     `json:"result"`
	ResultCount             int         `json:"resultCount"`
	PagedResultsCookie      interface{} `json:"pagedResultsCookie"`
	TotalPagedResultsPolicy string      `json:"totalPagedResultsPolicy"`
	TotalPagedResults       int         `json:"totalPagedResults"`
	RemainingPagedResults   int         `json:"remainingPagedResults"`
}

// NewClient -
func NewClient(host, xopenamusername, xopenampassword *string) (*Client, error) {
	c := Client{
		//TODO: this is just a hack for TLS verification only for testing. We should move to a configuration paramater for the provider
		HTTPClient: &http.Client{Timeout: 10 * time.Second, Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // test server certificate is not trusted.
			},
		},
		},
		// Default AM URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if (xopenamusername != nil) && (xopenampassword != nil) {

		// authenticate
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/am/json/realms/root/authenticate", c.HostURL), nil)
		if err != nil {
			return nil, err
		}

		// set forgerock Auth headers
		req.Header.Set("X-OpenAM-Username", *xopenamusername)
		req.Header.Set("X-OpenAM-Password", *xopenampassword)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept-API-Version", "resource=2.0, protocol=1.0")

		body, err := c.DoRequest(req)

		if err != nil {
			return nil, err
		}

		// parse response body
		ar := AuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}

		c.amadminSsotoken = ar.TokenID
	}

	return &c, nil
}

func (c *Client) DoRequest(req *http.Request) ([]byte, error) {
	if c.amadminSsotoken != "" {
		req.Header.Set("iplanetDirectoryPro", c.amadminSsotoken)
		req.Header.Set("Accept-API-Version", "resource=1.0, protocol=2.1")
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
