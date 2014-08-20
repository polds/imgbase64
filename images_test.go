package imgbase64

import (
	"testing"
)

func TestEncode(t *testing.T) {
	const in, out = "Hello, World!", "SGVsbG8sIFdvcmxkIQ=="
	if x := encode([]byte(in)); string(x) != out {
		t.Errorf("Encode(%v) = %s, want %v", in, x, out)
	}
}

func TestDefaultImage(t *testing.T) {
	const expected = "https://www.github.com/"

	SetDefaultImage(expected)
	if actual := DefaultImage(); actual != expected {
		t.Errorf("SetDefaultImage(%v) = %v, want %v", expected, actual, expected)
	}
}

func TestCleanUrl(t *testing.T) {
	const in, out = "Hello, World!", "Hello,%20World!"

	if x := cleanUrl(in); x != out {
		t.Errorf("CleanUrl(%v) = %v, want %v", in, x, out)
	}
}

func TestFromLocal(t *testing.T) {
	const expect = `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyRpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuMy1jMDExIDY2LjE0NTY2MSwgMjAxMi8wMi8wNi0xNDo1NjoyNyAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENTNiAoTWFjaW50b3NoKSIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDozNTQzMjU0NzM1MjIxMUUzQTMwNUJEOTg2QjYzRDA0NCIgeG1wTU06RG9jdW1lbnRJRD0ieG1wLmRpZDozNTQzMjU0ODM1MjIxMUUzQTMwNUJEOTg2QjYzRDA0NCI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOjM1NDMyNTQ1MzUyMjExRTNBMzA1QkQ5ODZCNjNEMDQ0IiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOjM1NDMyNTQ2MzUyMjExRTNBMzA1QkQ5ODZCNjNEMDQ0Ii8+IDwvcmRmOkRlc2NyaXB0aW9uPiA8L3JkZjpSREY+IDwveDp4bXBtZXRhPiA8P3hwYWNrZXQgZW5kPSJyIj8+3QswKwAAABBJREFUeNpi+P//PwNAgAEACPwC/tuiTRYAAAAASUVORK5CYII=`

	img, err := FromLocal("test.png")

	if err != nil {
		t.Error("There is an error retrieving the file")
	}

	if img != expect {
		t.Error("FromLocal(test.png) mismatch from expected output")
	}
}
