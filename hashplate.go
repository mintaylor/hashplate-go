package hashplatego

import (
	"strings"
)

type Options struct {
	CN        bool
	Emoji     bool
	ExcludeOI bool
	Seed      uint32
}

var defaultOptions = Options{
	Emoji: true,
	CN:    true,
}

func splitMix32(seed uint32) func() float64 {
	return func() float64 {
		seed += 0x9e3779b9
		z := seed
		z = (z ^ (z >> 16)) * 0x21f0aaad
		z = (z ^ (z >> 15)) * 0x735a2d97
		z ^= z >> 15
		return float64(z) / 4294967296
	}
}

func Hashplate(val string, opts Options) string {
	seed := opts.Seed
	if seed == 0 {
		for i := 0; i < len(val); i++ {
			seed = seed*31 + uint32(val[i])
		}
	}

	random := splitMix32(seed)

	randomStr := func(s string) byte {
		return s[int(random()*float64(len(s)))]
	}

	randomArr := func(arr []string) int {
		return int(random() * float64(len(arr)))
	}

	alphabet := Alphabet
	alphanum := Alphanum
	if opts.ExcludeOI {
		alphabet = AlphabetNoOI
		alphanum = AlphanumNoOI
	}

	var builder strings.Builder

	if opts.CN {
		builder.WriteRune(Province[int(random()*float64(len(Province)))])
		builder.WriteByte(randomStr(alphabet))
		builder.WriteString("·")
	}

	n := 7
	if opts.CN {
		n = 5
	}

	if !opts.CN && opts.Emoji {
		builder.WriteString(EmojiDict[randomArr(EmojiDict)])
		builder.WriteByte(' ')
	}

	for i := 0; i < n; i++ {
		if opts.CN {
			builder.WriteByte(randomStr(alphanum))
			continue
		}

		switch {
		case i <= 1:
			builder.WriteByte(randomStr(alphabet))
			if i == 1 {
				builder.WriteByte('-')
			}
		case i > 1 && i < 5:
			builder.WriteByte('0' + byte(random()*10))
			if i == 4 {
				builder.WriteByte('-')
			}
		default:
			builder.WriteByte(randomStr(alphabet))
		}

	}

	if !opts.CN && opts.Emoji {
		builder.WriteByte(' ')
		builder.WriteString(EmojiDict[randomArr(EmojiDict)])
	}

	str := builder.String()

	if opts.CN && opts.Emoji {
		prefix := EmojiDict[randomArr(EmojiDict)]
		suffix := EmojiDict[randomArr(EmojiDict)]
		return prefix + " " + str + " " + suffix
	}

	return str
}
