package report

import (
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../env.test")
}

func TestNewClient(t *testing.T) {
	c := NewClient("http://example.com", "abc123", 1)

	if c.ApiKey != "abc123" {
		t.Errorf("ApiKey should be abc123")
	}

	if c.HostUrl != "http://example.com" {
		t.Errorf("HostUrl should be http://example.com")
	}

	if c.SiteId != 1 {
		t.Errorf("SiteId should be 1")
	}
}

func TestGet(t *testing.T) {
	siteId, _ := strconv.Atoi(os.Getenv("MATOMO_SITE_ID"))

	c := NewClient(
		os.Getenv("MATOMO_HOST_URL"),
		os.Getenv("MATOMO_API_KEY"),
		siteId,
	)
	req := NewRequest()

	_, err := c.Get(req)

	if err != nil {
		t.Error(err)
		t.Errorf("Error should be nil")
	}
}
