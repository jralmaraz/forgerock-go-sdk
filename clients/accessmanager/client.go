package accessmanagerclient 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default environment URL
const HostURL string = "https://forgerock.iam.dtt-iam.xyz"

// Client -
type Client struct {
	HostURL         string
	HTTPClient      *http.Client
	amadminSsotoken string
}

// AuthStruct -
type AuthStruct struct {
	xopenamusername string `json:"xopenamusername"`
	xopenampassword string `json:"xopenampassword"`
}

// AuthResponse -
type AuthResponse struct {
	tokenID    string `json:"tokenId"`
	successURL string `json:"tokenId"`
	realm      string `json:"tokenId"`
}

// NewClient -
func NewClient(host, xopenamusername, xopenampassword *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default AM URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if (xopenamusername != nil) && (xopenampassword != nil) {

		// authenticate
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/am/json/realms/root/authenticate", c.HostURL), nil)
		req.Header.Set("X-OpenAM-Username", *xopenamusername)
		req.Header.Set("X-OpenAM-Password", *xopenampassword)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept-API-Version", "resource=2.0, protocol=1.0")
		if err != nil {
			return nil, err
		}

		body, err := c.doRequest(req)

		// parse response body
		ar := AuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}

		c.amadminSsotoken = ar.tokenID
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("iplanetDirectoryPro", c.amadminSsotoken)

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
