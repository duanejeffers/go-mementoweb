package timetravel

import (
	"testing"
	"time"
)

func TestGetMemento(t *testing.T) {
	var (
		err                           error
		url, expectedUrl, receivedUrl string
	)

	url = "https://en.wikipedia.org/wiki/Professional_wrestling"
	testTime := time.Date(2019, time.February, 14, 0, 0, 0, 0, time.UTC)

	expectedUrl = "http://wayback.archive-it.org/all/20190404150439/https://en.wikipedia.org/wiki/Professional_wrestling"

	receivedUrl, err = GetMemento(testTime, url)
	if err != nil {
		t.Error(err)
	}

	if receivedUrl != expectedUrl {
		t.Errorf("Url Mismatch. Expected: %s Received: %s", expectedUrl, receivedUrl)
	}
}
