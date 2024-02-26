package timetravel

import (
	"context"
	"fmt"
	"github.com/duanejeffers/go-mementoweb/httpclient"
	"net/http"
	"time"
)

const (
	BaseUrl = "http://timetravel.mementoweb.org"
	TimeFmt = "20060102150405"
)

func GetMemento(d time.Time, url string) (string, error) {
	return GetMementoCtx(context.Background(), d, url)
}

func GetMementoCtx(ctx context.Context, d time.Time, url string) (string, error) {
	var (
		timeStr string
	)

	// Determine time string:
	timeStr = d.Format(TimeFmt)

	return getMementoBase(ctx, fmt.Sprintf("%s/%s", timeStr, url))
}

func getMementoBase(ctx context.Context, timeAndURL string) (retUrl string, err error) {
	var (
		req *http.Request
		res *http.Response
		url string
	)
	url = fmt.Sprintf("%s/memento/%s", BaseUrl, timeAndURL)
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}

	res, err = httpclient.Client.Do(req)
	if err != nil && err != http.ErrUseLastResponse {
		return
	} else if res.StatusCode != 302 {
		return "", fmt.Errorf("GetMemento Error: MementoWeb responded with a non-302 status code. Code recieved: %d", res.StatusCode)
	}

	_ = res.Body.Close()

	retUrl = res.Header.Get("Location")
	return
}
