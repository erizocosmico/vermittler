package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Vermittler struct {
	Cfg *Config
}

func (v Vermittler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var img *Image
	var err error
	var format string

	r.ParseForm()
	filename := string(base64.URLEncoding.EncodeToString([]byte(r.Form.Encode())))

	imageExists := false
	if v.Cfg.CacheEnabled {
		files, err := ioutil.ReadDir(v.Cfg.CachePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, file := range files {
			fileParts := strings.Split(file.Name(), ".")
			if fileParts[0] == filename {
				imageExists = true
				format = fileParts[1]
				break
			}
		}
	}

	if !imageExists || !v.Cfg.CacheEnabled {
		img, err = NewImage(r.Form)
		err = img.Apply()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if format == "DS_Store" {
			return
		}
		img, err = NewImageFromFile(v.Cfg.CachePath + "/" + filename + "." + format)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if v.Cfg.CacheEnabled && !imageExists {
		go (func() {
			f, err := os.Create(v.Cfg.CachePath + "/" + filename + "." + img.Format)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = img.Write(f)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})()
	}

	w.Header().Add("Content-Type", "image/"+img.Format)
	err = img.Write(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	v := new(Vermittler)
	v.Cfg = NewConfig()
	http.ListenAndServe(":"+v.Cfg.Port, v)
}
