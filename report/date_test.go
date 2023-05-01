package report

import (
	"testing"
	"time"
)

func TestDateRangeLast(t *testing.T) {
	if DateRangeLast(1) != "last1" {
		t.Error("DateRangeLast(1) != last1")
	}
}

func TestDateRangePrevious(t *testing.T) {
	if DateRangePrevious(1) != "previous1" {
		t.Error("DateRangePrevious(1) != previous1")
	}
}

func TestDateRange(t *testing.T) {
	if DateRange(
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
		time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local),
	) != "2021-01-01,2021-01-02" {
		t.Error("DateRange is not correct")
	}
}
