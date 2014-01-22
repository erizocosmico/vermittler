vermittler
==========

Vermittler is an image processing server written in Go. It can scale, blur and cache images from the web and serve them.

Installation
-----
To install the latest version of vermittler (in master) you just have to get it with ```go get```

```bash
go get github.com/mvader/vermittler
```

Usage
----
To use vermittler you just have to run it passing the config file you want to use (config.json in your active directory by default)

```bash
vermittler -config /path/to/my/config.json
```

Config
----
Configuration parameters are very easy and straightforward.

```json
{
    "cache_enabled": true,
    "cache_path": "/path/to/my/dir",
    "port": "8888",
    "verbose": true
}
```
Requesting images from vermittler
----
To request an image from vermittler you have to send a GET HTTP request to a running instance of vermittler. For the following examples we'll asume that vermittler is running on http://localhost:8888.

There are 4 parameters that can be specified:
* **url**: the url of the image requested base64 encoded (not optional)
* **w**: width of the processed image (optional)
* **h**: height of the processed image (optional)
* **b**: blur radius (the higher the blurrier) (optional)

So if we want to request a blurried, 100x100 version of ```http://golang.org/doc/gopher/gopherbw.png``` we will use the following URL:
```
http://localhost:8888/?url=aHR0cDovL3d3dzIub3BlbnBob3RvLm5ldC92b2x1bWVzL3NpemVzL2tvcnJ5LzI1NTQzLzIuanBn&w=100&h=100&b=10
```
TODO
----
* More tests (test ```cache.go``` and ```vermittler.go```)
* More code coverage
* Add WebP format support.
