package browser

import (
	// standard
	"fmt"

	// 3rd-party
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

/* ---------------------------------------------------------------------------------------------- */

type Browser struct {
	launcher    *launcher.Launcher
	client      *rod.Browser
	currentPage *rod.Page
}

func New(width int, height int, pixelRatio float64) *Browser {
	l, c := getBrowser()

	c.DefaultDevice(devices.Device{
		Title: "LastFM WebP Widgets",
		Screen: devices.Screen{
			Vertical:         devices.ScreenSize{Width: width, Height: height},
			DevicePixelRatio: pixelRatio,
		},
	})

	c = c.MustConnect()

	return &Browser{l, c, nil}
}

/* ---------------------------------------------------------------------------------------------- */

func (b *Browser) GetClient() *rod.Browser {
	return b.client
}

func (b *Browser) GetCurrentPage() *rod.Page {
	return b.currentPage
}

func (b *Browser) Close() {
	b.launcher.Cleanup()
	b.launcher.Kill()
}

/* ---------------------------------------------------------------------------------------------- */

func (b *Browser) GoTo(url string) {
	b.currentPage = b.client.MustPage(url)
	b.currentPage.WaitNavigation(proto.PageLifecycleEventNameNetworkAlmostIdle)
}

func (b *Browser) TakeScreenshot(path string) error {
	if b.currentPage == nil {
		return fmt.Errorf("no page to take screenshot of")
	}

	b.currentPage.MustScreenshot(path)
	return nil
}
