package imgbase64

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
 * Default Image in case the requested image
 * does not exist.
 * @type {String}
 */
var df string = ""

/**
 * Get Default Image
 */
func DefaultImage() string {
	return df
}

/**
 * Set Default Image
 */
func SetDefaultImage(img string) {
	df = img
}

func encode(bin []byte) []byte {
	e64 := base64.StdEncoding

	maxEncLen := e64.EncodedLen(len(bin))
	encBuf := make([]byte, maxEncLen)

	e64.Encode(encBuf, bin)
	return encBuf[0:]
}

/**
 * Lightweight HTTP Client to fetch the image
 * Note: This will also pull webpages.
 * It is up to the user to use valid image urls.
 */
func get(url string) ([]byte, string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error getting url.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	ct := resp.Header.Get("Content-Type")

	if resp.StatusCode == 200 && len(body) > 512 {
		return body, ct
	}

	if DefaultImage() == "" {
		return []byte(""), ct
	}

	if url == DefaultImage() {
		panic("Catching an infinite loop! Default Image doesn't exist or is broken. Please rectify!")
	}

	return get(DefaultImage())
}

func NewImage(url string) string {
	image, ct := get(cleanUrl(url))
	enc := encode(image)

	switch ct {
	case "image/gif", "image/jpeg", "image/pjpeg", "image/png", "image/tiff":
		return fmt.Sprintf("data:%s;base64,%s", ct, enc)
	default:
	}

	return fmt.Sprintf("data:image/png;base64,%s", enc)
}

func cleanUrl(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
