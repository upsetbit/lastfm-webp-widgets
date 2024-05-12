package lastfm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type LastFmClientOptions struct {
	Username string
	APIKey   string
}

type LastFmClient struct {
	client   *http.Client
	endpoint string
	apikey   string
	Username string
}

func New(opts LastFmClientOptions) (*LastFmClient, error) {
	apikey := strings.TrimSpace(opts.APIKey)
	if apikey == "" {
		return nil, errors.New("API key is required")
	}

	client := LastFmClient{
		endpoint: "https://ws.audioscrobbler.com/2.0/",
		Username: opts.Username,
		apikey:   opts.APIKey,
		client:   &http.Client{},
	}

	return &client, nil
}

func (c *LastFmClient) buildRequest(method string) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "lastfm-webp-widgets")

	q := req.URL.Query()
	q.Add("method", method)
	q.Add("user", c.Username)
	q.Add("api_key", c.apikey)
	q.Add("format", "json")

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (c *LastFmClient) doAndHandleRequest(method string) ([]byte, error) {
	req, err := c.buildRequest(method)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *LastFmClient) GetUserInfo() (*LastFmUserInfo, error) {
	body, err := c.doAndHandleRequest("user.getinfo")
	if err != nil {
		return nil, err
	}

	var userInfo LastFmUserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func (c *LastFmClient) GetRecentTracks() (*LastFmUserRecentTracks, error) {
	body, err := c.doAndHandleRequest("user.getrecenttracks")
	if err != nil {
		return nil, err
	}

	var recentTracks LastFmUserRecentTracks
	err = json.Unmarshal(body, &recentTracks)
	if err != nil {
		return nil, err
	}

	return &recentTracks, nil
}
