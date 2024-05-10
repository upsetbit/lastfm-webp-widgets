package lastfm

import (
	"fmt"
	"os"

	"github.com/upsetbit/lastfm-webp-widgets/pkg/client/lastfm"
)

const (
	lastfmUsernameEnv = "LASTFM_USERNAME"
	lastfmAPIKeyEnv   = "LASTFM_API_KEY"
)

func BuildClient() (*lastfm.LastFMClient, error) {
	lastfmUsername, envIsSet := os.LookupEnv(lastfmUsernameEnv)
	if !envIsSet {
		return nil, fmt.Errorf("environment variable %s is required", lastfmUsernameEnv)
	}

	lastfmApiKey, envIsSet := os.LookupEnv(lastfmAPIKeyEnv)
	if !envIsSet {
		return nil, fmt.Errorf("environment variable %s is required", lastfmAPIKeyEnv)
	}

	lastfmClient, err := lastfm.New(lastfm.LastFMClientOptions{
		Username: lastfmUsername,
		APIKey:   lastfmApiKey,
	})

	if err != nil {
		return nil, err
	}

	return lastfmClient, nil
}
