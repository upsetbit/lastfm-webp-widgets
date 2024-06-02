package storage

import "bytes"

func Init() {
	storageInit()
}

func Save(key string, buf bytes.Buffer) {
	storageSave(key, buf)
}
