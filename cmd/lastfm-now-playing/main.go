package main

import (
	"embed"

	"github.com/upsetbit/lastfm-webp-widgets/internal/webputil"
)

//go:embed assets/images/test.webp
var assets embed.FS

func main() {
	testWebP, _ := assets.ReadFile("assets/images/test.webp")
	webputil.DoSomething(testWebP)
}
