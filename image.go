package main

import (
	"code.google.com/p/graphics-go/graphics"
	"encoding/base64"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Image represents an image and its values and transformations.
type Image struct {
	Width, Height int64
	Blur          float64
	URL           *url.URL
	Format        string
	Data          image.Image
}

// ValidFormat checks if the format is valid or not.
// TODO: Add webp format support
func ValidFormat(format string) bool {
	return format == "png" || format == "jpeg" || format == "gif"
}

func (i *Image) readImage(r io.Reader, remote bool) error {
	var data image.Image
	var err error

	switch i.Format {
	case "png":
		data, err = png.Decode(r)
		break

	case "jpeg":
		data, err = jpeg.Decode(r)
		break

	case "gif":
		data, err = gif.Decode(r)
		break
	}
	if err != nil {
		return err
	}

	i.Data = data

	if remote {
		if i.Width < 1 {
			i.Width = int64(i.Data.Bounds().Max.X)
		}

		if i.Height < 1 {
			i.Height = int64(i.Data.Bounds().Max.Y)
		}
	}

	return nil
}

// NewImage returns the image data for the specified parameters in the query string.
func NewImage(values url.Values) (*Image, error) {
	var w, h, b, u string
	var width, height int64
	var blur float64

	w = values.Get("w")
	h = values.Get("h")
	b = values.Get("b")
	u = values.Get("url")

	if u == "" {
		return nil, errors.New("url not provided")
	}

	URLBytes, err := base64.StdEncoding.DecodeString(u)
	if err != nil {
		return nil, errors.New("provided url wasn't properly base64 encoded")
	}

	parsedURL, err := url.Parse(string(URLBytes))
	if err != nil {
		return nil, errors.New("provided url is not valid")
	}

	img := &Image{
		Width:  -1,
		Height: -1,
		Blur:   -1.0,
		URL:    parsedURL,
		Format: "",
	}

	if w != "" {
		width, err = strconv.ParseInt(w, 10, 32)
		if err != nil {
			return nil, errors.New("invalid width parameter provided")
		}
		img.Width = width
	}

	if h != "" {
		height, err = strconv.ParseInt(h, 10, 32)
		if err != nil {
			return nil, errors.New("invalid height parameter provided")
		}
		img.Height = height
	}

	if b != "" {
		blur, err = strconv.ParseFloat(b, 32)
		if err != nil {
			return nil, errors.New("invalid blur radius parameter provided")
		}
		img.Blur = blur
	}

	res, err := http.Get(img.URL.String())
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	defer res.Body.Close()

	formatStrings := strings.Split(res.Header.Get("Content-Type"), "/")
	if len(formatStrings) != 2 && formatStrings[0] != "image" && ValidFormat(formatStrings[1]) {
		return nil, errors.New("invalid format. Expecting `image/jpeg`, `image/png` or `image/gif`")
	}
	img.Format = formatStrings[1]

	err = img.readImage(res.Body, true)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// Blurries the image.
func (i *Image) blur() error {
	dst := image.NewRGBA(i.Data.Bounds())
	err := graphics.Blur(dst, i.Data, &graphics.BlurOptions{StdDev: i.Blur})
	if err == nil {
		i.Data = dst
	}

	return err
}

// Scales the image.
func (i *Image) scale() error {
	dst := image.NewRGBA(image.Rect(0, 0, int(i.Width), int(i.Height)))
	err := graphics.Scale(dst, i.Data)
	if err == nil {
		i.Data = dst
	}

	return err
}

// Apply applies the needed transformations.
func (i *Image) Apply() error {
	var err error

	if i.Width > 0 && i.Width > 0 {
		err = i.scale()
		if err != nil {
			return err
		}
	}

	if i.Blur > 0 {
		err = i.blur()
		if err != nil {
			return err
		}
	}

	return nil
}

// Writes the transformed image to the specified writer.
func (i *Image) Write(w io.Writer) error {
	var err error

	switch i.Format {
	case "png":
		err = png.Encode(w, i.Data)
		break

	case "jpeg":
		err = jpeg.Encode(w, i.Data, nil)
		break

	case "gif":
		err = gif.Encode(w, i.Data, nil)
		break
	}
	return err
}
