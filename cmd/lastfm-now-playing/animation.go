package main

import (
	// standard
	"bytes"
	"os"

	// internal
	"github.com/upsetbit/lastfm-webp-widgets/internal/util"
	"github.com/upsetbit/lastfm-webp-widgets/pkg/webpanimation"
)

func animate(frames []string, output string) {
	wpa := webpanimation.NewWebpAnimation(WIDGET_WIDTH*WIDGET_PIXEL_RATIO, WIDGET_HEIGHT*WIDGET_PIXEL_RATIO, 0)

	// see: <https://developers.google.com/speed/webp/docs/gif2webp#:~:text=%2Dkmin%20int%2C%20%2Dkmax%20int>
	wpa.WebPAnimEncoderOptions.SetKmin(9)
	wpa.WebPAnimEncoderOptions.SetKmax(17)

	defer wpa.ReleaseMemory() // don't forget call this or you will have memory leaks

	webpConfig := webpanimation.NewWebpConfig()
	webpConfig.SetLossless(1)

	timeline := 0
	for _, f := range frames {
		frame, err := util.LoadImageFromFS(f)
		if err != nil {
			panic(err)
		}

		err = wpa.AddFrame(frame, timeline, webpConfig)
		if err != nil {
			panic(err)
		}

		timeline += 200
	}

	err := wpa.AddFrame(nil, timeline, webpConfig)
	if err != nil {
		panic(err)
	}

	log.Info("encoding animation...")
	var buf bytes.Buffer
	err = wpa.Encode(&buf)
	if err != nil {
		panic(err)
	}

	os.WriteFile(output, buf.Bytes(), 0644)
	log.Info("webp animation saved", "output", output)
}
