package util

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
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
