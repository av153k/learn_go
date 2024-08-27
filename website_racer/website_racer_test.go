package websiteracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	t.Run("Compare speeds of racers - returning the fastest one", func(t *testing.T) {
		slowServer := makeTestServer(20 * time.Millisecond)
		fastServer := makeTestServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Client()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := WebsiteRacer(slowUrl, fastUrl)

		if err != nil {
			t.Errorf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if server doesn't respond within specified time", func(t *testing.T) {
		server := makeTestServer(25 * time.Second)

		defer server.Close()

		_, err := ConfigurableWebsiteRacer(server.URL, server.URL, 20*time.Second)

		if err == nil {
			t.Errorf("Expected and error but got nil")
		}
	})

}

func makeTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
