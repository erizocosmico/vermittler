package main

import "net/http"

type Vermittler struct{}

func (v Vermittler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    
    img, err := NewImage(r.Form)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
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
    // TODO: Configurable port
    http.ListenAndServe(":8888", v)
}
