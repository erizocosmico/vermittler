vermittler
==========

Vermittler is an image processing server written in Go. It can scale, blur and cache images from the web and serve them.

Installation
-----
To install the latest version of vermittler (in master) you just have to get it with ```go get```

```
go get github.com/mvader/vermittler
```

Usage
----
To use vermittler you just have to run it passing the config file you want to use (config.json in your active directory by default)

```
vermittler -config /path/to/my/config.json
```

Config
----
Configuration parameters are very easy and straightforward.

```
{
    "cache_enabled": true,
    "cache_path": "/path/to/my/dir",
    "port": "8888",
    "verbose": true
}
```
TODO
----
* More tests (test ```cache.go``` and ```vermittler.go```)
* More code coverage
* Add WebP format support.