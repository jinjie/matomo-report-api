# Go Matomo Reporting API

Simple implementation of Matomo Reporting API

There are still a lot to work on...

## Install 

`go get github.com/jinjie/matomo-report-api`

## Usage

```go
package main

import (
	"fmt"

	"github.com/jinjie/matomo-report-api/report"
)

func main() {
	c := report.NewClient(
		"https://matomo.example.com/",
		"yourapikey",
		1, // idSite
	)

	req := report.NewRequest()

	req.Date = report.DATE_YESTERDAY

	result, err := c.Get(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

```
