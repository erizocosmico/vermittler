package vermittler

import (
	"net/url"
	"testing"
    "os"
)

func values() url.Values {
	query := "url=aHR0cDovL2dvbGFuZy5vcmcvZG9jL2dvcGhlci9nb3BoZXJidy5wbmc=&w=150&h=200&b=10"
    
	values, err := url.ParseQuery(query)
	if err != nil {
		panic("query `" + query + "` is not a valid query string.")
	}
    
    return values
}

func TestNewImageimg(t *testing.T) {
    imageUrl := "http://golang.org/doc/gopher/gopherbw.png"
	values := values()

	img, err := NewImage(values)
	if err != nil {
		t.Fatal(err)
	}

	if img.URL.String() != imageUrl {
		t.Errorf("Expecting url `%s` to be `%s`", img.URL.String(), imageUrl)
	}

	if img.Width != 150 {
		t.Errorf("Expecting width `%d` to be `%d`", img.Width, 150)
	}

	if img.Height != 200 {
		t.Errorf("Expecting height `%d` to be `%d`", img.Height, 200)
	}

	if img.Blur != 10 {
		t.Errorf("Expecting blur `%d` to be `%d`", img.Blur, 10)
	}

	if img.Format != "png" {
		t.Errorf("Expecting format `%s` to be `%s`", img.Format, "png")
	}
    
    // TODO: Test .Data somehow
}

func TestWrite(t *testing.T) {
    values := values()
    
	img, err := NewImage(values)
	if err != nil {
		t.Fatal(err)
	}
    
    f, err := os.Create("test."+img.Format)
    if err != nil {
        t.Fatal(err)
    }
    
    err = img.Write(f)
    if err != nil {
        t.Fatal(err)
    }
    
    err = os.Remove("test."+img.Format)
    if err != nil {
        t.Fatal(err)
    }
}

func TestBlur(t *testing.T) {
    // TODO: Implement
}

func TestScale(t *testing.T) {
    // TODO: Implement
}