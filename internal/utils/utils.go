package utils

import (
	"log"
	"os"
	"path/filepath"
)

var ManifestBinaryPath, _ = filepath.Abs("AndroidManifest.xml")
var OldAppNameNormalized string

func WriteChanges(raw []byte, path string) {
	//Open a new file for writing only
	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write bytes to file
	_, err = file.Write(raw)
	if err != nil {
		log.Panic("Failed to write changes to disk", err)
	}
}

func isValidFile(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
