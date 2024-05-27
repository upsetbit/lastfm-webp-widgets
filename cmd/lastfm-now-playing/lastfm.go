package main

import (
	// standard
	"fmt"
	"os"

	// internal
	"github.com/upsetbit/lastfm-webp-widgets/internal/util"
	L "github.com/upsetbit/lastfm-webp-widgets/pkg/client/lastfm"
)

const (
	env_LastFmUsername = "LASTFM_USERNAME"
	env_LastFmAPIKey   = "LASTFM_API_KEY"
)

var (
	lastfm *L.LastFmClient
)

func initLastFmClient() {
	lastfmUsername, envIsSet := os.LookupEnv(env_LastFmUsername)
	if !envIsSet {
		panic(fmt.Sprintf("unset env %s", env_LastFmUsername))
	}

	lastfmApiKey, envIsSet := os.LookupEnv(env_LastFmAPIKey)
	if !envIsSet {
		panic(fmt.Sprintf("unset env %s", env_LastFmAPIKey))
	}

	lastfmClient, err := L.New(L.LastFmClientOptions{
		Username: lastfmUsername,
		APIKey:   lastfmApiKey,
	})

	if err != nil {
		panic(err)
	}

	lastfm = lastfmClient
	log.Info("lastfm client created", "username", lastfm.Username)
}

func getLastFmUserInfo() *L.LastFmUser {
	data, err := lastfm.GetUserInfo()
	if err != nil {
		panic(err)
	}

	d := data.User
	log.Info(
		"got authenticated user info",
		"username",
		d.Name,
		"realname",
		d.RealName,
		"scrobbles",
		d.PlayCount,
	)

	return &d
}

func getLastFmUserRecentTracks() *L.LastFmRecentTracks {
	data, err := lastfm.GetRecentTracks()
	if err != nil {
		panic(err)
	}

	d := data.RecentTracks
	if len(d.Track) == 0 {
		log.Warn("user has not listened to any tracks yet")
		log.Info("nothing to do; exiting")
		return nil
	}

	lastTrack := d.Track[0]
	log.Info(
		"got recent tracks",
		"artist_name",
		lastTrack.Artist.Text,
		"is_playing",
		util.BoolToYesOrNo(lastTrack.Attr.NowPlaying == "true"),
	)

	return &d
}
