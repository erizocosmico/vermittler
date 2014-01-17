package main

import (
	"os"
	"testing"
    "fmt"
)

type ImageTest struct {
    Width, Height int64
    Blur float64
    URL, Format, Query string
}

var (
    imageTest1 = ImageTest{
        Width: 150,
        Height: 200,
        Blur: 10.0,
        URL: "http://golang.org/doc/gopher/gopherbw.png",
        Format: "png",
        Query: "url=aHR0cDovL2dvbGFuZy5vcmcvZG9jL2dvcGhlci9nb3BoZXJidy5wbmc=&w=150&h=200&b=10",
    }
    imageTest2 = ImageTest{
        Width: 400,
        Height: 400,
        Blur: -1.0,
        URL: "http://www2.openphoto.net/volumes/sizes/korry/25543/2.jpg",
        Format: "jpeg",
        Query: "url=aHR0cDovL3d3dzIub3BlbnBob3RvLm5ldC92b2x1bWVzL3NpemVzL2tvcnJ5LzI1NTQzLzIuanBn&w=400&h=400",
    }
    imageTest3 = ImageTest{
        Width: 802,
        Height: 610,
        Blur: -1.0,
        URL: "http://www2.openphoto.net/volumes/sizes/korry/25543/2.jpg",
        Format: "jpeg",
        Query: "url=aHR0cDovL3d3dzIub3BlbnBob3RvLm5ldC92b2x1bWVzL3NpemVzL2tvcnJ5LzI1NTQzLzIuanBn",
    }
    imageTest4 = ImageTest{
        Width: 200,
        Height: 100,
        Blur: -1.0,
        URL: "http://media3.giphy.com/media/13yyvZdx0W6kTK/giphy.gif",
        Format: "gif",
        Query: "url=aHR0cDovL21lZGlhMy5naXBoeS5jb20vbWVkaWEvMTN5eXZaZHgwVzZrVEsvZ2lwaHkuZ2lm&w=200&h=100",
    }
)

func remoteImageTest(testImg ImageTest, t *testing.T) {
    values := parseQueryString(testImg.Query)
	img, err := NewImage(values)
	if err != nil {
		t.Fatal(err)
	}
    
    err = img.Apply()
	if err != nil {
		t.Fatal(err)
	}

	if img.URL.String() != testImg.URL {
		t.Errorf("Expecting url `%s` to be `%s`", img.URL.String(), testImg.URL)
	}

    if testImg.Width > 0 {
    	if img.Width != testImg.Width {
    		t.Errorf("Expecting width `%d` to be `%d`", img.Width, testImg.Width)
    	}
    }

	if testImg.Height > 0 {
    	if img.Height != testImg.Height {
    		t.Errorf("Expecting height `%d` to be `%d`", img.Height, testImg.Height)
    	}
    }

	if testImg.Blur > 0 {
    	if img.Blur != testImg.Blur {
    		t.Errorf("Expecting blur `%f` to be `%f`", img.Blur, testImg.Blur)
    	}
    }

	if img.Format != testImg.Format {
		t.Errorf("Expecting format `%s` to be `%s`", img.Format, testImg.Format)
	}
}

func TestNewImage(t *testing.T) {
    fmt.Println("Testing remote: "+imageTest1.URL)
	remoteImageTest(imageTest1, t)
    
    fmt.Println("Testing remote: "+imageTest2.URL)
    remoteImageTest(imageTest2, t)
    
    fmt.Println("Testing remote: "+imageTest3.URL)
    remoteImageTest(imageTest3, t)
    
    fmt.Println("Testing remote: "+imageTest4.URL)
    remoteImageTest(imageTest4, t)
}

func TestWrite(t *testing.T) {
	values := parseQueryString(imageTest1.Query)

	img, err := NewImage(values)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create("test." + img.Format)
	if err != nil {
		t.Fatal(err)
	}

	err = img.Write(f)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove("test." + img.Format)
	if err != nil {
		t.Fatal(err)
	}
}