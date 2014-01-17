package main

import (
	"io/ioutil"
	"os"
	"strings"
)

// Returns if the file is in the cache path and if it is also returns the format of the image
func FileInCache(filename, cachePath string) (bool, string, error) {
	files, err := ioutil.ReadDir(cachePath)
	if err != nil {
		return false, "", err
	}

	for _, file := range files {
		fileParts := strings.Split(file.Name(), ".")
		if fileParts[0] == filename {
			return true, fileParts[1], nil
		}
	}

	return false, "", nil
}

// Loads an image from the cache
func LoadImage(filename, cachePath string) (*Image, error) {
	var err error

	img := &Image{
		Width:  -1,
		Height: -1,
		Blur:   -1.0,
	}

	f, err := os.Open(cachePath+"/"+filename)
	if err != nil {
		return nil, err
	}
    defer f.Close()
    
	formatStrings := strings.Split(filename, ".")
	img.Format = formatStrings[len(formatStrings)-1]
    
    err = img.readImage(f, false)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// Stores the file in the cache path
func CacheFile(file string, img *Image) {
	f, err := os.Create(file)
	if err == nil {
		defer f.Close()
		img.Write(f)
	}
}
