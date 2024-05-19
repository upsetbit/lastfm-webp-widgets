package util

import (
	_ "image/png"

	"bytes"
	"embed"
	"image"
	"os"
)

func LoadFileFromEmbedFS(fs embed.FS, path string) *bytes.Reader {
	file, _ := fs.ReadFile(path)
	return bytes.NewReader(file)
}

func LoadImageFromEmbedFS(fs embed.FS, path string) (image.Image, error) {
	fileReader := LoadFileFromEmbedFS(fs, path)
	data, _, err := image.Decode(fileReader)
	return data, err
}

func LoadImageFromFS(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, _, err := image.Decode(file)
	return data, err
}
