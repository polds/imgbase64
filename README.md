# imgbase64

[![Build Status](https://drone.io/github.com/polds/imgbase64/status.png)](https://drone.io/github.com/polds/imgbase64/latest)

Convert an image to its base64 equivalent. [API Documentation](http://godoc.org/github.com/polds/imgbase64) on go.pkgdoc.org.

## [Features](https://github.com/polds/imgbase64#features)

* Fetch an image from an external server and convert to base64
* Use an image on your local machine to convert to base64
* Share a `bytes.Buffer` to the package to detect the mime type and convert the image to base64
* Ability to set a failover image if desired image throws an error or does not exist.
* Automatically converts urls with " " to %20. **Note:** If your target domain uses +, _, etc. for space indicators it is your responsibility to modify the url accordingly.
* Correctly sets the base64 encoding type to the media type of the image. Example: `data:image/jpeg;base64`
* Returned image is ready to be used


## [Installation](https://github.com/polds/imgbase64#installation)

```
go get github.com/polds/imgbase64
```

## [Usage](https://github.com/polds/imgbase64#usage)



```golang
imgbase64.SetDefaultImage("http://yourdomain.com/defaultimage.png") // Optional - If FromRemote fails it will return " "

// Fetch the image from a remote server
img := imgbase64.FromRemote("http://somedomain.com/animage.jpg")

// Use an image local to machine
img := imgbase64.FromLocal("test.png")

// Pass an image buffer to package
img := imgbase64.FromBuffer(b)

```

## [Notes](https://github.com/polds/imgbase64#notes)

* Please make sure your Default Image is a working image url. If your Default Image fails to load the package will `panic()`. This was deliberate.
* `FromLocal` has several instances where it might `panic()`, these are deliberate.

## [Todo](https://github.com/polds/imgbase64#todo)

* Better error detection. Websites with custom 404 pages will sometimes "succeed" as being images and the package will base64 the html content of that page.
* ~~Ability to use local images to convert to base64~~
* Cache the Default Image only once during the lifetime of the application. Presently every failed image is repulled.
* Deprecate `NewImage`
* Allow use of DefaultImage with `FromBuffer` and `FromLocal`

## [Change Log](https://github.com/polds/imgbase64#change-log)

##### October 24, 2013
 * ADD: FromLocal - Allows you to base64 encode local images
 * ADD: FromBuffer - Allows you to base64 encode from a `byte.Buffer`
 * CHANGE: Rename function `NewImage` to `FromRemote`, please update accordingly. Will remove `NewImage` in future commits.
 * ADD: Examples



This is my first package in Go. Suggestions or comments please tweet [@Peter_Olds](https://twitter.com/Peter_Olds)