package lastfm

type LastFmImageMetadata struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}

// -------------------------------------------------------------------------------------------------

type LastFmUserRegistrationDate struct {
	Unixtime string `json:"unixtime"`
	Text     int64  `json:"#text"`
}

type LastFmUser struct {
	Name        string                     `json:"name"`
	Age         string                     `json:"age"`
	Subscriber  string                     `json:"subscriber"`
	RealName    string                     `json:"realname"`
	Bootstrap   string                     `json:"bootstrap"`
	PlayCount   string                     `json:"playcount"`
	ArtistCount string                     `json:"artist_count"`
	Playlists   string                     `json:"playlists"`
	TrackCount  string                     `json:"track_count"`
	AlbumCount  string                     `json:"album_count"`
	Image       []LastFmImageMetadata      `json:"image"`
	Registered  LastFmUserRegistrationDate `json:"registered"`
	Country     string                     `json:"country"`
	Gender      string                     `json:"gender"`
	URL         string                     `json:"url"`
	Type        string                     `json:"type"`
}

type LastFmUserInfo struct {
	User LastFmUser `json:"user"`
}

// -------------------------------------------------------------------------------------------------

type LastFmScrobbleTimestamp struct {
	UTS  string `json:"uts"`
	Text string `json:"#text"`
}

type LastFmMusicIdentification struct {
	MBID string `json:"mbid"`
	Text string `json:"#text"`
}

type LastFmUserAttr struct {
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	Total      string `json:"total"`
	TotalPages string `json:"totalPages"`
	User       string `json:"user"`
}

type LastFmTrackAttr struct {
	NowPlaying string `json:"nowplaying"`
}

type LastFmTrack struct {
	Name       string                    `json:"name"`
	Album      LastFmMusicIdentification `json:"album"`
	Artist     LastFmMusicIdentification `json:"artist"`
	Image      []LastFmImageMetadata     `json:"image"`
	MBID       string                    `json:"mbid"`
	Streamable string                    `json:"streamable"`
	URL        string                    `json:"url"`
	Date       LastFmScrobbleTimestamp   `json:"date"`
	Attr       LastFmTrackAttr           `json:"@attr"`
}

type LastFmRecentTracks struct {
	Attr  LastFmUserAttr `json:"@attr"`
	Track []LastFmTrack  `json:"track"`
}

type LastFmUserRecentTracks struct {
	RecentTracks LastFmRecentTracks `json:"recenttracks"`
}
