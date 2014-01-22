package main

import (
	"flag"
	"net/http"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.json", "config file for the application")
	flag.Parse()

	v := new(Vermittler)
	v.Cfg = NewConfig(configFile)

	http.ListenAndServe(":"+v.Cfg.Port, v)
}
