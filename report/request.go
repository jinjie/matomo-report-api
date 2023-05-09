package report

// See https://developer.matomo.org/api-reference/reporting-api for detailed
// information about the parameters.
type Request struct {
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
	METHOD_ACTIONS_GET                                   = "Actions.get"
	METHOD_ACTIONS_GET_PAGE_URLS                         = "Actions.getPageUrls"
	METHOD_ACTIONS_GET_PAGE_URLS_FOLLOWING_SITE_SEARCH   = "Actions.getPageUrlsFollowingSiteSearch"
	METHOD_ACTIONS_GET_PAGE_TITLES_FOLLOWING_SITE_SEARCH = "Actions.getPageTitlesFollowingSiteSearch"
	METHOD_ACTIONS_GET_PAGE_TITLES                       = "Actions.getPageTitles"
	METHOD_ACTIONS_GET_DOWNLOADS                         = "Actions.getDownloads"
	METHOD_ACTIONS_GET_OUTLINKS                          = "Actions.getOutlinks"

	METHOD_VISITS_SUMMARY_GET                  = "VisitsSummary.get"
	METHOD_VISITS_SUMMARY_GET_UNIQUE_VISITORS  = "VisitsSummary.getUniqueVisitors"
	METHOD_VISITS_SUMMARY_GET_VISITS           = "VisitsSummary.getVisits"
	METHOD_VISITS_SUMMARY_GET_ACTIONS          = "VisitsSummary.getActions"
	METHOD_VISITS_SUMMARY_GET_MAX_ACTIONS      = "VisitsSummary.getMaxActions"
	METHOD_VISITS_SUMMARY_GET_BOUNCE_COUNT     = "VisitsSummary.getBounceCount"
	METHOD_VISITS_SUMMARY_GET_VISITS_CONVERTED = "VisitsSummary.getVisitsConverted"

	METHOD_EVENTS_GET_CATEGORY                = "Events.getCategory"
	METHOD_EVENTS_GET_ACTION                  = "Events.getAction"
	METHOD_EVENTS_GET_NAME                    = "Events.getName"
	METHOD_EVENTS_GET_ACTION_FROM_CATEGORY_ID = "Events.getActionFromCategoryId"
	METHOD_EVENTS_GET_NAME_FROM_ACTION_ID     = "Events.getNameFromActionId"
	METHOD_EVENTS_GET_CATEGORY_FROM_NAME_ID   = "Events.getCategoryFromNameId"

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
