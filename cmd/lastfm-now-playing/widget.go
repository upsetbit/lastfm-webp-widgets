package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getWidgetLocation() (string, string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	widgetPath := filepath.Join(wd, "assets", "widget-now-playing.html")
	widgetFile, err := os.Stat(widgetPath)
	if err != nil {
		panic(err)
	}

	if widgetFile.IsDir() {
		panic(fmt.Errorf("path %s is a directory", widgetPath))
	}

	if widgetFile.Size() == 0 {
		panic(fmt.Errorf("file %s is empty", widgetPath))
	}

	return wd, widgetPath
}
