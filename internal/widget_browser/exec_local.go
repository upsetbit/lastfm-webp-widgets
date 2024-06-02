//go:build exec_local

package browser

import (
	"os"

	// 3rd-party
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func getBrowser() (*launcher.Launcher, *rod.Browser) {
	var u *launcher.Launcher

	if binp, binpIsSet := os.LookupEnv("CHROMIUM_BROWSER_BINARY_PATH"); binpIsSet {
		u = launcher.New().Bin(binp)
	} else {
		u = launcher.New()
	}

	return u, rod.New().ControlURL(u.MustLaunch())
}
