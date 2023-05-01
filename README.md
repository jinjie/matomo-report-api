# Go Matomo Reporting API

Simple implementation of Matomo Reporting API

There are still a lot to work on...

[Matomo Reporting API Documentation](https://developer.matomo.org/api-reference/reporting-api)

## Install 

`go get github.com/jinjie/matomo-report-api`

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/jinjie/matomo-report-api/report"
)

func main() {
	c := report.NewClient(
		"https://demo.matomo.cloud/",
		"anonymous",
		1,
	)

	// Get the number of visits yesterday
	req := report.NewRequest().
		SetPeriod(report.PERIOD_DAY).
		SetDate("yesterday")

	yesterday := &report.ActionsGetMetric{}
	err := c.Get(req, yesterday)
	if err != nil {
		panic(err)
	}

	// Print the result
	fmt.Println("Number of visits yesterday:")
	fmt.Printf("%#v\n", yesterday)

	// Get the number of visits for the last 30 days
	req = report.NewRequest().
		SetPeriod(report.PERIOD_DAY).
		SetDate("last30")

	last30 := make(map[string]*report.ActionsGetMetric)
	err = c.Get(req, &last30)

	if err != nil {
		panic(err)
	}

	// Print the result
	fmt.Println("Number of visits last 30 days:")
	for i, v := range last30 {
		fmt.Printf("%s: %#v\n", i, v)
	}

	// Get the number of visits between 2 dates
	from := time.Date(2023, 3, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(2023, 4, 15, 0, 0, 0, 0, time.Local)
	req = report.NewRequest().
		SetPeriod(report.PERIOD_DAY).
		SetDate(
			report.DateRange(
				from,
				to,
			),
		)

	dateRange := make(map[string]*report.ActionsGetMetric)
	err = c.Get(req, &dateRange)

	if err != nil {
		panic(err)
	}

	// Print the result
	fmt.Printf(
		"Number of visits between %s and %s:\n",
		from.Format("2006-01-02"),
		to.Format("2006-01-02"),
	)

	for i, v := range dateRange {
		fmt.Printf("%s: %#v\n", i, v)
	}
}
```
