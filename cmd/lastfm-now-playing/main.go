package main

import (
	_ "github.com/joho/godotenv/autoload"

	// standard
	"os"
	"strconv"

	// internal
	"github.com/upsetbit/lastfm-webp-widgets/internal/util"
)

func main() {
	log.Info("program started")
	titleNeedsScroll := false

	framedir, err := os.MkdirTemp("", "lastfm-now-playing-frames-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(framedir)

	_, widgetPath := getWidgetLocation()
	initBrowser(widgetPath)
	initLastFmClient()
	initS3Client()

	lastFmUser := getLastFmUserInfo()
	lastFmUserRecentTracks := getLastFmUserRecentTracks()
	if lastFmUserRecentTracks == nil {
		return
	}
	lastTrack := lastFmUserRecentTracks.Track[0]

	setUserURL(lastFmUser.Name)
	setUserStats(lastFmUser.PlayCount, util.UnixToHumanReadable(lastFmUser.Registered.Text))
	setTrackTitle(lastTrack.Name)
	setArtistName(lastTrack.Artist.Text)
	setAlbumCoverSource(lastTrack.Image[len(lastTrack.Image)-1].Text)

	if getTrackTitleSizeInPixels() > TRACK_TITLE_MAX_SIZE_PIXELS {
		log.Info("size is too long, making it scrollable")
		setTrackTitleScrollable(lastTrack.Name)
		titleNeedsScroll = true
	}

	if lastTrack.Attr.NowPlaying == "true" {
		setUserListeningNow()
	} else {
		timeSince, err := strconv.ParseInt(lastTrack.Date.UTS, 10, 64)
		if err != nil {
			panic(err)
		}
		setUserListeningLastPlayed(util.UnixToRelativeHumanTime(timeSince))
	}

	waitPageToLoad()

	if titleNeedsScroll {
		animateWithScroll(framedir, getScrollTicksAmount())
	} else {
		animateSimple(framedir)
	}

	log.Info("program finished")
}
