package format_test

import (
	"testing"

	"github.com/BF-Moritz/utils.go/format"
)

func TestRound(t *testing.T) {
	rounded := format.Round(float32(3.212361), 2)
	if rounded != "3.21" {
		t.Fail()
	}

	rounded = format.Round(1236312.0, 2)
	if rounded != "1236312" {
		t.Fail()
	}
}

func TestRoundToSignificant(t *testing.T) {
	rounded := format.RoundToSignificant(float32(3.212361), 2)
	if rounded != "3.2" {
		t.Logf("'%s' is not equal the expected '%s'\n", rounded, "3.2")
		t.Fail()
	}

	rounded = format.RoundToSignificant(1236312.0, 4)
	if rounded != "1236000" {
		t.Logf("'%s' is not equal the expected '%s'\n", rounded, "1236000")
		t.Fail()
	}
}
