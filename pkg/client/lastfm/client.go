package lastfm

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type LastFMClientOptions struct {
	Username string
	APIKey   string
}

type LastFMClient struct {
	client   *http.Client
	endpoint string
	apikey   string
	username string
}

func New(opts LastFMClientOptions) (*LastFMClient, error) {
	apikey := strings.TrimSpace(opts.APIKey)
	if apikey == "" {
		return nil, errors.New("API key is required")
	}

	client := LastFMClient{
		endpoint: "https://ws.audioscrobbler.com/2.0/",
		username: opts.Username,
		apikey:   opts.APIKey,
		client:   &http.Client{},
	}

	return &client, nil
}

func (c *LastFMClient) buildRequest(method string) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "lastfm-webp-widgets")

	q := req.URL.Query()
	q.Add("method", method)
	q.Add("user", c.username)
	q.Add("api_key", c.apikey)
	q.Add("format", "json")

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (c *LastFMClient) GetRecentTracks() (*LastFMUserRecentTracks, error) {
	req, err := c.buildRequest("user.getrecenttracks")
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("Response: %s\n", string(body))

	var recentTracks LastFMUserRecentTracks
	err = json.Unmarshal(body, &recentTracks)
	if err != nil {
		return nil, err
	}

	return &recentTracks, nil
}
