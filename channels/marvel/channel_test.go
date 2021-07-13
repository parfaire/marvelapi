package marvel

import (
	"net/url"
	"strings"
	"testing"
)

func TestAppendGETAuthentication(t *testing.T) {
	targetString := "www.google.com"
	targetURL, _ := url.Parse(targetString)

	c := Channel{}
	c.AppendGETAuthentication(targetURL)

	if targetURL.String() == targetString {
		t.Errorf("AppendGETAuthentication does not work")
	}
	if targetURL.String()[:len(targetString)] != targetString {
		t.Errorf("Got = %v | Expected =%v", targetURL.String()[:len(targetString)], targetString)
	}
	containsApikey := strings.Contains(targetURL.String(), "apikey=")
	containsTs := strings.Contains(targetURL.String(), "ts=")
	containsHash := strings.Contains(targetURL.String(), "hash=")
	if !containsTs || !containsApikey || !containsHash {
		t.Errorf("Missing appended parameter")
	}
}

func TestAppendOffsetLimit(t *testing.T) {
	targetString := "www.google.com"
	targetURL, _ := url.Parse(targetString)

	offset := 5
	limit := 10
	c := Channel{}
	c.AppendOffsetLimit(targetURL, offset, limit)

	if targetURL.String() == targetString {
		t.Errorf("AppendOffsetLimit does not work")
	}
	if targetURL.String()[:len(targetString)] != targetString {
		t.Errorf("Got = %v | Expected =%v", targetURL.String()[:len(targetString)], targetString)
	}
	containsOffset := strings.Contains(targetURL.String(), "offset=")
	containsLimit := strings.Contains(targetURL.String(), "limit=")
	if !containsOffset || !containsLimit {
		t.Errorf("Missing appended parameter")
	}
}
