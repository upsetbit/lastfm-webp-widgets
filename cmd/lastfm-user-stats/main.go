package main

import (
	_ "github.com/joho/godotenv/autoload"

	"fmt"

	"github.com/upsetbit/lastfm-webp-widgets/internal/lastfm"
)

func main() {
	lastfmClient, err := lastfm.BuildClient()
	if err != nil {
		panic(err)
	}

	res, err := lastfmClient.GetRecentTracks()
	if err != nil {
		panic(err)
	}

	recentTracks := res.RecentTracks
	fmt.Printf("user: %s\n", recentTracks.Attr.User)
	fmt.Printf("found %s tracks\n\n", recentTracks.Attr.Total)

	if len(recentTracks.Track) == 0 {
		return
	}

	lastListenedTrack := recentTracks.Track[0]
	if lastListenedTrack.Attr.NowPlaying == "true" {
		fmt.Printf("now playing: %s\n", lastListenedTrack.Name)
	} else {
		fmt.Printf("last listened: %s\n", lastListenedTrack.Name)
	}
}
