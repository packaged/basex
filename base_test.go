package basex

import (
	"log"
	"math"
	"testing"
	"time"
)

var b2 = []rune("01")
var b8 = []rune("01234567")
var b10 = []rune("0123456789")
var b16 = []rune("0123456789ABCDEF")
var b32 = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUV")
var b36 = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var b62 = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var b64 = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
var humanBase32 = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")
var urlSafe = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789$-_.!*()")
var extended = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@$%^&*()[]{}:;,<>")

var bases = []struct {
	baseRune []rune
	bitSize  int
	base     int
}{
	{b2, 1, 2},
	{b8, 3, 8},
	{b10, 4, 10},
	{b16, 4, 16},
	{b32, 5, 32},
	{b36, 6, 36},
	{b62, 6, 62},
	{b64, 6, 64},
	{humanBase32, 5, 32},
	{urlSafe, 7, 70},
	{extended, 7, 80},
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
