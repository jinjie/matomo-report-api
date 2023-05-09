package report

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

type Requestable interface {
	QueryString() (q string, err error)
}

// See https://developer.matomo.org/api-reference/reporting-api for detailed
// information about the parameters.
type Request struct {
	Requestable

	Module string `url:"module"` // eg, API
	Method string `url:"method"` // eg, VisitsSummary.getVisits

	SiteId      int    `url:"idSite"`                 // eg, 1 or 1,2,3
	Period      string `url:"period"`                 // one of PERIOD_*
	Date        string `url:"date"`                   // one of DATE_* or range of date. eg, 2014-01-01,2014-01-31
	Format      string `url:"format"`                 // one of FORMAT_*
	FilterLimit int    `url:"filter_limit,omitempty"` // eg, 10

	// eg, country==United States
	// @see https://developer.matomo.org/api-reference/reporting-api-segmentation
	Segment string `url:"segment,omitempty"`

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

func NewRequest() *Request {
	return &Request{
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

func (r *Request) SetSiteId(id int) *Request {
	r.SiteId = id

	return r
}

func (r *Request) SetPeriod(period string) *Request {
	r.Period = period

	return r
}

func (r *Request) SetDate(date string) *Request {
	r.Date = date

	return r
}

func (r *Request) SetModule(module string) *Request {
	r.Module = module

	return r
}

type BulkRequest struct {
	Requestable

	Module    string    `url:"module"` // eg, API
	Method    string    `url:"method"` // eg, VisitsSummary.getVisits
	Format    string    `url:"format"` // one of FORMAT_*
	Requests  []Request `url:"-"`
	TokenAuth string    `url:"token_auth"` // eg, abc123
}

func NewBulkRequest() *BulkRequest {
	return &BulkRequest{
		Module: "API",
		Method: "API.getBulkRequest",
		Format: FORMAT_JSON,
	}
}

func (r *BulkRequest) QueryString() (q string, err error) {
	v, err := query.Values(r)

	if err != nil {
		return "", err
	}

	q = v.Encode()

	for i, req := range r.Requests {
		qs, _ := req.QueryString()
		q += fmt.Sprintf("&urls[%d]=%s", i, qs)
	}

	return q, err
}
