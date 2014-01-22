package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// FileInCache returns if the file is in the cache path and if it is also returns the format of the image
func FileInCache(filename string, cfg *Config) (bool, string, error) {
	files, err := ioutil.ReadDir(cfg.CachePath)
	if err != nil {
		return false, "", err
	}

	for _, file := range files {
		fileParts := strings.Split(file.Name(), ".")
		if fileParts[0] == filename {
			if cfg.Verbose {
				fmt.Println("vermittler: file `" + filename + "." + fileParts[1] + "` exists in cache")
			}

			return true, fileParts[1], nil
		}
	}

	return false, "", nil
}

// LoadImage loads an image from the cache
func LoadImage(filename string, cfg *Config) (*Image, error) {
	var err error

	img := &Image{
		Width:  -1,
		Height: -1,
		Blur:   -1.0,
	}

	f, err := os.Open(cfg.CachePath + "/" + filename)
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

	if cfg.Verbose {
		fmt.Println("vermittler: loading image `" + filename + "` from cache")
	}

	return img, nil
}

// CacheFile stores the file in the cache path
func CacheFile(file string, img *Image, cfg *Config) {
	f, err := os.Create(cfg.CachePath + "/" + file)
	if err == nil {
		defer f.Close()
		img.Write(f)

		if cfg.Verbose {
			fmt.Println("vermittler: storing file `" + file + "` in cache")
		}
	} else {
		panic(err.Error())
	}
}
