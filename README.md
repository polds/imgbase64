# imgbase64

### This branch is not passing build! Please proceed carefully.


[![Build Status](https://drone.io/github.com/polds/imgbase64/status.png)](https://drone.io/github.com/polds/imgbase64/latest)

Convert an image to its base64 equivalent. [API Documentation](http://godoc.org/github.com/polds/imgbase64) on go.pkgdoc.org.

# Features

* Ability to set a failover image if desired image throws an error or does not exist.
* Automatically converts urls with " " to %20. **Note:** If your target domain uses +, _, etc. for space indicators it is your responsibility to modify the url accordingly.
* Correctly sets the base64 encoding type to the media type of the image. Example: `data:image/jpeg;base64`
* Returned image is ready to be used


# Installation

```
go get github.com/polds/imgbase64
```

# Usage


```golang
imgbase64.SetDefaultImage("http://yourdomain.com/defaultimage.png") // Optional - If NewImage fails it will return " "
url := imgbase64.NewImage("http://somedomain.com/animage.jpeg")
```

# Notes

* Please make sure your Default Image is a working image url. If your Default Image fails to load the package will `panic()`. This was deliberate.

# Todo

* Better error detection. Websites with custom 404 pages will sometimes "succeed" as being images and the package will base64 the html content of that page.
* Ability to use local images to convert to base64
* Cache the Default Image only once during the lifetime of the application. Presently every failed image is repulled.

This is my first package in Go. Suggestions or comments please tweet [@Peter_Olds](https://twitter.com/Peter_Olds)