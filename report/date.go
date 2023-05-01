package report

import (
	"fmt"
	"time"
)

func DateRangeLast(days int) string {
	return fmt.Sprintf("last%d", days)
}

func DateRangePrevious(days int) string {
	return fmt.Sprintf("previous%d", days)
}

func DateRange(from time.Time, to time.Time) string {
	return fmt.Sprintf(
		"%s,%s",
		from.Format("2006-01-02"),
		to.Format("2006-01-02"),
	)
}
