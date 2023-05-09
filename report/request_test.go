package report

import (
	"fmt"
	"testing"
)

func TestBulkRequestQueryString(t *testing.T) {
	bulkReq := NewBulkRequest()

	req1 := NewRequest().SetSiteId(1)
	req2 := NewRequest().SetSiteId(2)
	bulkReq.Requests = append(bulkReq.Requests, *req1)
	bulkReq.Requests = append(bulkReq.Requests, *req2)

	q, err := bulkReq.QueryString()
	if err != nil {
		t.Errorf("error should be nil")
	}

	fmt.Printf("q: %s\n", q)
}
