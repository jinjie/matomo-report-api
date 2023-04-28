package report

type ActionsGetMetric struct {
	PageViews       int `json:"nb_pageviews"`
	UniquePageViews int `json:"nb_uniq_pageviews"`
	Downloads       int `json:"nb_downloads"`
	UniqueDownloads int `json:"nb_uniq_downloads"`
	Outlinks        int `json:"nb_outlinks"`
	UniqueOutlinks  int `json:"nb_uniq_outlinks"`
	Searches        int `json:"nb_searches"`
	Keywords        int `json:"nb_keywords"`
}
