package report

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Metric interface{}

type Client struct {
	ApiKey  string
	HostUrl string
	SiteId  int
}

func NewClient(HostUrl string, ApiKey string, SiteId int) *Client {
	return &Client{
		ApiKey:  ApiKey,
		HostUrl: HostUrl,
		SiteId:  SiteId,
	}
}

func (c *Client) BulkExecute(req *BulkRequest) (json string, err error) {
	return "", err
}

func (c *Client) Execute(req *Request) (json string, err error) {
	q, err := req.QueryString()
	if err != nil {
		return "", err
	}

	resp, err := http.Get(
		fmt.Sprintf(
			"%s?%s",
			c.HostUrl,
			q,
		),
	)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { // other than 200
		return "", fmt.Errorf("status code: %d", resp.StatusCode)
	}

	resultBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(resultBytes), nil
}

// Get data and return Metric
func (c *Client) Get(req Request, m interface{}) (err error) {
	// todo design a better way to check if req is empty
	if req.Method == "" {
		return errors.New("use NewRequest() to create a request")
	}

	if req.SiteId == 0 { // if not set, client default
		req.SiteId = c.SiteId
	}

	req.TokenAuth = c.ApiKey

	// Will always return JSON format for now
	req.Format = FORMAT_JSON

	j, err := c.Execute(req)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(j), &m)
	if err != nil {
		return err
	}

	return nil
}
