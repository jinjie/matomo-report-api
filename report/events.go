package report

type EventsGetMetric struct {
	Label                  string `json:"label"`
	Visits                 int    `json:"nb_visits"`
	Events                 int    `json:"nb_events"`
	EventsWithValue        int    `json:"nb_events_with_value"`
	SumEventValue          int    `json:"sum_event_value"`
	MinEventValue          int    `json:"min_event_value"`
	MaxEventValue          int    `json:"max_event_value"`
	SumDailyNbUniqVisitors int    `json:"sum_daily_nb_uniq_visitors"`
	AvgEventValue          int    `json:"avg_event_value"`
	Segment                string `json:"segment"`
	Idsubdatatable         int    `json:"idsubdatatable"`
}
