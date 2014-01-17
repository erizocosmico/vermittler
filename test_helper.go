package main

import (
    "net/url"
)

func parseQueryString(query string) url.Values {
	values, err := url.ParseQuery(query)
	if err != nil {
		panic("query `" + query + "` is not a valid query string.")
	}

	return values
}