package strutil

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/width"
)

// RunePos alias of the strings.IndexRune
func RunePos(s string, ru rune) int {
	return strings.IndexRune(s, ru)
}

// IsSpaceRune returns true if the given rune is a space, otherwise false.
func IsSpaceRune(r rune) bool {
	return r <= 256 && IsSpace(byte(r)) || unicode.IsSpace(r)
}

// Utf8Len of the string
func Utf8Len(s string) int { return utf8.RuneCountInString(s) }

// Utf8len of the string
func Utf8len(s string) int { return utf8.RuneCountInString(s) }

// RuneCount of the string
func RuneCount(s string) int { return len([]rune(s)) }

// RuneWidth of the rune.
//
// Example:
// 	RuneWidth('你') // 2
// 	RuneWidth('a') // 1
// 	RuneWidth('\n') // 0
func RuneWidth(r rune) int {
	p := width.LookupRune(r)
	k := p.Kind()

	// eg: "\n"
	if k == width.Neutral {
		return 0
	}

	if k == width.EastAsianFullwidth || k == width.EastAsianWide || k == width.EastAsianAmbiguous {
		return 2
	}
	return 1
}

// TextWidth alias of the Utf8Width()
func TextWidth(s string) int { return Utf8Width(s) }

// Utf8Width utf8 string width.
//
//	str := "hi,你好"
//	strutil.Utf8Width(str)	=> 7
//	len(str) => 9
//	len([]rune(str)) = utf8.RuneCountInString(s) => 5
func Utf8Width(s string) (size int) {
	for _, runeVal := range []rune(s) {
		size += RuneWidth(runeVal)
	}
	return size
}

// TextTruncate alias of the Utf8Truncate()
func TextTruncate(s string, w int, tail string) string { return Utf8Truncate(s, w, tail) }

// Utf8Truncate a string with given width.
func Utf8Truncate(s string, w int, tail string) string {
	if sw := Utf8Width(s); sw <= w {
		return s
	}

	i := 0
	r := []rune(s)
	w -= TextWidth(tail)

	tmpW := 0
	for ; i < len(r); i++ {
		cw := RuneWidth(r[i])
		if tmpW+cw > w {
			break
		}
		tmpW += cw
	}
	return string(r[0:i]) + tail
}

// TextSplit alias of the Utf8Split()
func TextSplit(s string, w int) []string { return Utf8Split(s, w) }

// Utf8Split split a string by width.
func Utf8Split(s string, w int) (ss []string) {
	if sw := Utf8Width(s); sw <= w {
		return []string{s}
	}

	tmpW := 0
	tmpS := ""
	for _, r := range []rune(s) {
		rw := RuneWidth(r)
		if tmpW+rw == w {
			tmpS += string(r)
			ss = append(ss, tmpS)

			tmpW, tmpS = 0, "" // reset
			continue
		}

		if tmpW+rw > w {
			ss = append(ss, tmpS)

			// append to next line.
			tmpW, tmpS = rw, string(r)
			continue
		}

		tmpW += rw
		tmpS += string(r)
	}

	if tmpW > 0 {
		ss = append(ss, tmpS)
	}
	return
}

// TextWrap a string by "\n"
func TextWrap(s string, w int) string { return WidthWrap(s, w) }

// WidthWrap a string by "\n"
func WidthWrap(s string, w int) string {
	tmpW := 0
	out := ""

	for _, r := range []rune(s) {
		cw := RuneWidth(r)
		if r == '\n' {
			out += string(r)
			tmpW = 0
			continue
		} else if tmpW+cw > w {
			out += "\n"
			tmpW = 0
			out += string(r)
			tmpW += cw
			continue
		}

		out += string(r)
		tmpW += cw
	}
	return out
}
