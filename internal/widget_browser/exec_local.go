//go:build exec_local

package browser

import (
	// 3rd-party
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func getBrowser() (*launcher.Launcher, *rod.Browser) {
	u := launcher.New()
	return u, rod.New().ControlURL(u.MustLaunch())
}
