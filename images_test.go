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
