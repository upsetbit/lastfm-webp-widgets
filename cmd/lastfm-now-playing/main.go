package main

import (
	"bytes"
	"embed"
	"os"
	"path/filepath"

	"github.com/upsetbit/lastfm-webp-widgets/internal/util"
	"github.com/upsetbit/lastfm-webp-widgets/internal/webpanimation"
)

//go:embed assets/images/*.png
var assets embed.FS

func main() {
	pngFrames := []string{"frame1.png", "frame2.png", "frame3.png", "frame4.png", "frame5.png"}
	timeline := 0

	var buf bytes.Buffer
	var err error

	webpanim := webpanimation.NewWebpAnimation(1062, 938, 0)

	// see: <https://developers.google.com/speed/webp/docs/gif2webp#:~:text=%2Dkmin%20int%2C%20%2Dkmax%20int>
	webpanim.WebPAnimEncoderOptions.SetKmin(9)
	webpanim.WebPAnimEncoderOptions.SetKmax(17)

	defer webpanim.ReleaseMemory() // don't forget call this or you will have memory leaks

	webpConfig := webpanimation.NewWebpConfig()
	webpConfig.SetLossless(1)

	for _, f := range pngFrames {
		frame, err := util.LoadImageFromEmbedFS(assets, filepath.Join("assets/images", f))
		if err != nil {
			panic(err)
		}

		err = webpanim.AddFrame(frame, timeline, webpConfig)
		if err != nil {
			panic(err)
		}

		timeline += 200
	}

	err = webpanim.AddFrame(nil, timeline, webpConfig)
	if err != nil {
		panic(err)
	}

	err = webpanim.Encode(&buf) // encode animation and write result bytes in buffer
	if err != nil {
		panic(err)
	}

	os.WriteFile("animation.webp", buf.Bytes(), 0777) // write bytes on disk
}
