package websiteracer

import (
	"errors"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func WebsiteRacer(url1, url2 string) (string, error) {
	return ConfigurableWebsiteRacer(url1, url2, tenSecondTimeout)
}

func ConfigurableWebsiteRacer(url1, url2 string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", errors.New("server timeout")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
