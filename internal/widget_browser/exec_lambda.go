//go:build exec_lambda

package browser

import (
	// 3rd-party
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func getBrowser() (*launcher.Launcher, *rod.Browser) {
	u := launcher.
		New().
		Bin("/opt/google/chrome/chrome").
		Set("allow-running-insecure-content").
		Set("autoplay-policy", "user-gesture-required").
		Set("disable-component-update").
		Set("disable-domain-reliability").
		Set("disable-features", "AudioServiceOutOfProcess", "IsolateOrigins", "site-per-process").
		Set("disable-print-preview").
		Set("disable-setuid-sandbox").
		Set("disable-site-isolation-trials").
		Set("disable-speech-api").
		Set("disable-web-security").
		Set("disk-cache-size", "33554432").
		Set("enable-features", "SharedArrayBuffer").
		Set("hide-scrollbars").
		Set("ignore-gpu-blocklist").
		Set("in-process-gpu").
		Set("mute-audio").
		Set("no-default-browser-check").
		Set("no-pings").
		Set("no-sandbox").
		Set("no-zygote").
		Set("single-process").
		Set("use-gl", "swiftshader")

	return u, rod.New().ControlURL(u.MustLaunch())
}
