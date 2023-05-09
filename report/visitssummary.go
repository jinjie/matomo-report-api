package report

type VisitSummaryGetMetric struct {
	UniqueVisitors  int     `json:"nb_uniq_visitors"`
	Users           int     `json:"nb_users"`
	Visits          int     `json:"nb_visits"`
	Actions         int     `json:"nb_actions"`
	VisitsConverted int     `json:"nb_visits_converted"`
	BounceCount     int     `json:"bounce_count"`
	SumVisitLength  int     `json:"sum_visit_length"`
	MaxActions      int     `json:"max_actions"`
	BounceRate      string  `json:"bounce_rate"`
	ActionsPerVisit float64 `json:"nb_actions_per_visit"`
	AvgTimeOnSite   int     `json:"avg_time_on_site"`
}
