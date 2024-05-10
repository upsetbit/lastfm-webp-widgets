package main

import (
	"embed"

	"github.com/upsetbit/lastfm-webp-widgets/internal/webputil"
)

//go:embed assets/images/a.webp
var assets embed.FS

func main() {
	testWebP, _ := assets.ReadFile("assets/images/a.webp")
	webputil.DoSomething(testWebP)
}
