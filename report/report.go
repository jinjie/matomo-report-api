package report

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
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

func (c *Client) GetJSON(req Request) (json string, err error) {
	req.SiteId = c.SiteId
	req.TokenAuth = c.ApiKey

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

	j, err := c.GetJSON(req)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(j), &m)
	if err != nil {
		return err
	}

	return nil
}

// See https://developer.matomo.org/api-reference/reporting-api for detailed
// information about the parameters.
type Request struct {
	Module string `url:"module"` // eg, API
	Method string `url:"method"` // eg, VisitsSummary.getVisits

	SiteId int    `url:"idSite"` // eg, 1 or 1,2,3
	Period string `url:"period"` // one of PERIOD_*
	Date   string `url:"date"`   // one of DATE_* or range of date. eg, 2014-01-01,2014-01-31
	Format string `url:"format"` // one of FORMAT_*

	TokenAuth string `url:"token_auth"` // eg, abc123
}

const (
	PERIOD_DAY   = "day"
	PERIOD_WEEK  = "week"
	PERIOD_MONTH = "month"
	PERIOD_YEAR  = "year"
	PERIOD_RANGE = "range"

	DATE_TODAY      = "today"
	DATE_YESTERDAY  = "yesterday"
	DATE_LAST_WEEK  = "lastWeek"
	DATE_LAST_MONTH = "lastMonth"
	DATE_LAST_YEAR  = "lastYear"

	FORMAT_XML      = "xml"
	FORMAT_JSON     = "json"
	FORMAT_CSV      = "csv"
	FORMAT_TSV      = "tsv"
	FORMAT_HTML     = "html"
	FORMAT_RSS      = "rss"
	FORMAT_ORIGINAL = "original"
)

func NewRequest() Request {
	return Request{
		Module: "API",
		Method: "Actions.get",
		Period: PERIOD_DAY,
		Date:   DATE_TODAY,
		Format: FORMAT_JSON,
	}
}

func (r *Request) QueryString() (q string, err error) {
	v, err := query.Values(r)
	return v.Encode(), err
}
