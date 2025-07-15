package basex

import (
	"errors"
	"math"
	"math/bits"
)

var ErrInvalidCharacter = errors.New("invalid character in base encoding")

type Base struct {
	base         int
	bitSize      int
	alpha        []rune
	baseMap      map[rune]int
	replacements map[rune]rune
	padding      rune
}

func NewBase(alpha []rune) *Base {
	b := &Base{
		base:         len(alpha),
		alpha:        alpha,
		baseMap:      make(map[rune]int),
		replacements: make(map[rune]rune),
	}

	b.bitSize = bits.Len(uint(b.base) - 1)
	if b.bitSize == 0 {
		b.bitSize = 1
	} else if b.bitSize > 8 {
		panic("invalid base size, must be <= 256")
	}

	for i := 0; i < b.base; i++ {
		b.baseMap[b.alpha[i]] = i
	}

	return b
}

func (b Base) mapReplacements(input string) string {
	if len(b.replacements) == 0 {
		return input
	}

	var result []rune
	for _, char := range input {
		if replacement, ok := b.replacements[char]; ok {
			result = append(result, replacement)
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

func (b Base) EncodeInt(value uint64) string {
	if value == 0 {
		return string(b.alpha[0])
	}

	var result []rune
	for value > 0 {
		remainder := value % uint64(b.base)
		result = append([]rune{b.alpha[remainder]}, result...)
		value /= uint64(b.base)
	}

	return string(result)
}

func (b Base) DecodeInt(encoded string) (uint64, error) {
	if len(encoded) == 0 {
		return 0, nil
	}

	var value uint64
	encoded = b.mapReplacements(encoded)
	length := len(encoded) - 1
	for pos, char := range encoded {
		charPos := length - pos
		if index, ok := b.baseMap[char]; ok {
			value += uint64(index) * uint64(math.Pow(float64(b.base), float64(charPos)))
		} else {
			return 0, ErrInvalidCharacter
		}
	}
	return value, nil
}
