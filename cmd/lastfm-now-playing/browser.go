package main

import (
	// standard
	"fmt"
	"strings"
	"time"

	// internal
	widgetBrowser "github.com/upsetbit/lastfm-webp-widgets/internal/widget_browser"
)

var browser *widgetBrowser.Browser

func initBrowser(fp string) {
	browser = widgetBrowser.New(WIDGET_WIDTH, WIDGET_HEIGHT, WIDGET_PIXEL_RATIO)
	log.Info("browser initialized")

	widgetURI := fmt.Sprintf("file://%s", fp)
	browser.GoTo(widgetURI)
	log.Info("went to widget location", "uri", widgetURI)
}

/* ---------------------------------------------------------------------------------------------- */

func getTrackTitleSizeInPixels() int {
	size := browser.
		GetCurrentPage().
		MustEval("() => music.getTrackTitleSizeInPixelsRounded()").
		Int()

	log.Info("got track title size in pixels", "size", size)
	return size
}

func setTrackTitle(title string) {
	tt := browser.
		GetCurrentPage().
		MustEval("(t) => music.setTrackTitle(t)", title).
		Str()

	log.Info("track title setted", "title", tt)
}

func setTrackTitleScrollable(title string) {
	whitespace := strings.Repeat("&nbsp;", TRACK_TITLE_SCROLLABLE_WHITESPACES)
	titleScrollable := fmt.Sprintf("%s%s%s", title, whitespace, title)

	setTrackTitle(titleScrollable)
}

func setArtistName(name string) {
	artist := browser.
		GetCurrentPage().
		MustEval("(n) => music.setArtistName(n)", name).
		Str()

	log.Info("artist name setted", "name", artist)
}

func setAlbumCoverSource(src string) {
	cover := browser.
		GetCurrentPage().
		MustEval("(s) => music.setAlbumCoverSource(s)", src).
		Str()

	log.Info("album cover source setted", "src", cover)
}

/* ---------------------------------------------------------------------------------------------- */

func setUserURL(username string) {
	url := browser.
		GetCurrentPage().
		MustEval("(u) => user.setUrl(u)", username).
		Str()

	log.Info("user URL setted", "url", url)
}

func setUserStats(scrobbles string, creation string) {
	stats := browser.
		GetCurrentPage().
		MustEval("(s, c) => user.setStats(s, c)", scrobbles, creation).
		Str()

	log.Info("user stats setted", "stats", stats)
}

func setUserListeningNow() {
	status := browser.
		GetCurrentPage().
		MustEval("() => user.setListeningStatusNowPlaying()").
		Str()

	log.Info("user listening status setted", "status", status)
}

func setUserListeningLastPlayed(relativeTime string) {
	status := browser.
		GetCurrentPage().
		MustEval("(t) => user.setListeningStatusLastPlayed(t)", relativeTime).
		Str()

	log.Info("user listening status setted", "status", status)
}

/* ---------------------------------------------------------------------------------------------- */

func randomizeSoundWave() {
	browser.
		GetCurrentPage().
		MustEval("() => waves.randomize()")
}

/* ---------------------------------------------------------------------------------------------- */

func waitPageToLoad() {
	log.Info("waiting for page to load")

	start := time.Now()
	browser.GetCurrentPage().MustWaitStable()
	diffMS := time.Since(start).Milliseconds()

	log.Info("page loaded", "time", diffMS)
}
