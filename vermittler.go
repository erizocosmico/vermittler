package main

import (
	"encoding/base64"
	"net/http"
)

// Vermittler handler
type Vermittler struct {
	Cfg *Config
}

// Serves the app
func (v Vermittler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var img *Image
	var err error
	var format string

	r.ParseForm()
	filename := string(base64.URLEncoding.EncodeToString([]byte(r.Form.Encode())))

	imageExists := false
	if v.Cfg.CacheEnabled {
		imageExists, format, err = FileInCache(filename, v.Cfg.CachePath)
	}

	if !imageExists || !v.Cfg.CacheEnabled {
		img, err = NewImage(r.Form)
		err = img.Apply()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// For some reason OS X does weird things with .DS_Store
		if format == "DS_Store" {
			return
		}
		img, err = LoadImage(filename + "." + format, v.Cfg.CachePath)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if v.Cfg.CacheEnabled && !imageExists {
		go CacheFile(v.Cfg.CachePath+"/"+filename+"."+img.Format, img)
	}

	w.Header().Add("Content-Type", "image/"+img.Format)
	err = img.Write(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}