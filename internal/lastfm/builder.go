package lastfm

import (
	"fmt"
	"os"

	"github.com/upsetbit/lastfm-webp-widgets/pkg/client/lastfm"
)

const (
	env_LastFmUsername = "LASTFM_USERNAME"
	env_LastFmAPIKey   = "LASTFM_API_KEY"
)

func BuildClient() (*lastfm.LastFmClient, error) {
	lastfmUsername, envIsSet := os.LookupEnv(env_LastFmUsername)
	if !envIsSet {
		return nil, fmt.Errorf("unset env %s", env_LastFmUsername)
	}

	lastfmApiKey, envIsSet := os.LookupEnv(env_LastFmAPIKey)
	if !envIsSet {
		return nil, fmt.Errorf("unset env %s", env_LastFmAPIKey)
	}

	lastfmClient, err := lastfm.New(lastfm.LastFmClientOptions{
		Username: lastfmUsername,
		APIKey:   lastfmApiKey,
	})

	if err != nil {
		return nil, err
	}

	return lastfmClient, nil
}
