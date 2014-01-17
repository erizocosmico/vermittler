package main

import "net/http"

func main() {
	v := new(Vermittler)
	v.Cfg = NewConfig()
	http.ListenAndServe(":"+v.Cfg.Port, v)
}
