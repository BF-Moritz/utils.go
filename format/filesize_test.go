package format_test

import (
	"testing"

	"github.com/BF-Moritz/utils.go/format"
)

func TestFSBase10(t *testing.T) {
	tests := []struct {
		f string
		e string
	}{
		{
			f: format.FSBase10(213263126735),
			e: "213GB",
		},
		{
			f: format.FSBase10(21326312673),
			e: "21.3GB",
		},
		{
			f: format.FSBase10(2132631267),
			e: "2.13GB",
		},
		{
			f: format.FSBase10(2132631263264537735),
			e: "2.13EB",
		},
		{
			f: format.FSBase10(2136123),
			e: "2.14MB",
		},
		{
			f: format.FSBase10(1),
			e: "1B",
		},
	}

	for _, test := range tests {
		if test.e != test.f {
			t.Logf("'%s' is not equal the expected '%s'\n", test.f, test.e)
			t.Fail()
		}
	}
}
