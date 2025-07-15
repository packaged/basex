package basex

import (
	"log"
	"math"
	"testing"
	"time"
)

var bases = []struct {
	baseRune []rune
	bitSize  int
	base     int
}{
	{B2, 1, 2},
	{B8, 3, 8},
	{B10, 4, 10},
	{B16, 4, 16},
	{B32, 5, 32},
	{B36, 6, 36},
	{B62, 6, 62},
	{B64, 6, 64},
	{Human32, 5, 32},
	{UrlSafe, 7, 70},
	{Extended, 7, 80},
}

func TestNewBase(t *testing.T) {
	for _, test := range bases {
		b := NewBase(test.baseRune)
		if b.base != test.base {
			t.Errorf("Expected base %d, got %d for base %s", test.base, b.base, string(test.baseRune))
		}
		if b.bitSize != test.bitSize {
			t.Errorf("Expected bit size %d, got %d for base %s", test.bitSize, b.bitSize, string(test.baseRune))
		}
		if len(b.alpha) != len(test.baseRune) {
			t.Errorf("Expected alpha length %d, got %d for base %s", len(test.baseRune), len(b.alpha), string(test.baseRune))
		}
	}
}

func TestEncodeInt(t *testing.T) {

	inputs := []uint64{0, 1, 2, 10, 16, 32, 36, 62, 63, 64, 255, 256, 100, 1000, 10000, 1000000, uint64(time.Now().Unix()), math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64}
	for _, test := range bases {
		b := NewBase(test.baseRune)
		for _, input := range inputs {
			enc := b.EncodeInt(input)
			dec, err := b.DecodeInt(enc)
			log.Println(enc, dec)
			if err != nil {
				t.Errorf("Error decoding value %s for base %s: %v", enc, string(test.baseRune), err)
				continue
			}
			if dec != input {
				t.Errorf("Expected encoded value %d, got %d for base %s", input, dec, string(test.baseRune))
				continue
			}
		}
	}
}
