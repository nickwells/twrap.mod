package twrap

import (
	"fmt"
	"math"
	"strings"
)

// WrapPrefixed will print the text as with Wrap. The first line will start
// with the prefix and the indent of the subsequent lines will be adjusted to
// include the length of the prefix.
func (twc TWConf) WrapPrefixed(prefix, text string, indent int) {
	twc.Wrap3Indent(prefix+text,
		indent, indent+len(prefix), indent+len(prefix))
}

// Wrap will print the text onto the configured writer but wraps and indents
// the text. It will split the text into paragraphs at any newline
// characters. It will always print at least MinCharsToPrint chars but will
// try to fit the text into TargetLineLen chars.
func (twc TWConf) Wrap(text string, indent int) {
	twc.Wrap3Indent(text, indent, indent, indent)
}

// Wrap2Indent will print the text as with Wrap. The first line printed of
// each paragraph will be indented by the first supplied indent and the other
// lines will be indented by the second indent value.
func (twc TWConf) Wrap2Indent(
	text string,
	firstLineIndent, otherLineIndent int,
) {
	twc.Wrap3Indent(text, firstLineIndent, firstLineIndent, otherLineIndent)
}

// Wrap3Indent will print the text as with Wrap. The first line printed will
// be indented by the first supplied indent, thereafter the first line of
// each paragraph will be indented by the second supplied indent and any
// other lines will be indented by the third indent value.
func (twc TWConf) Wrap3Indent(
	text string,
	firstLineIndent, paraFirstLineIndent, otherLineIndent int,
) {
	firstLineMaxWidth := int(
		math.Max(
			float64(twc.MinCharsToPrint),
			float64(twc.TargetLineLen-firstLineIndent)))
	paraFirstLineMaxWidth := int(
		math.Max(
			float64(twc.MinCharsToPrint),
			float64(twc.TargetLineLen-paraFirstLineIndent)))
	otherLineMaxWidth := int(
		math.Max(
			float64(twc.MinCharsToPrint),
			float64(twc.TargetLineLen-otherLineIndent)))

	paras := strings.Split(text, "\n")
	prefix := strings.Repeat(" ", otherLineIndent)
	firstLinePrefix := strings.Repeat(" ", firstLineIndent)
	maxWidth := firstLineMaxWidth

	for _, para := range paras {
		fmt.Fprint(twc.W, firstLinePrefix)

		lineLen := 0
		word := make([]rune, 0, len(para))
		spaces := make([]rune, 0, maxWidth)
		for _, r := range para {
			if r == ' ' {
				if len(word) > 0 {
					if lineLen == 0 {
						lineLen = len(word)
						word = twc.printAndClear(word)
						spaces = spaces[:0]
					} else if lineLen+len(word)+len(spaces) <= maxWidth {
						lineLen += len(word) + len(spaces)
						spaces = twc.printAndClear(spaces)
						word = twc.printAndClear(word)
					} else {
						fmt.Fprintln(twc.W)
						maxWidth = otherLineMaxWidth
						fmt.Fprint(twc.W, prefix)
						lineLen = len(word)
						word = twc.printAndClear(word)
						spaces = spaces[:0]
					}
				}

				spaces = append(spaces, r)
			} else {
				word = append(word, r)
			}
		}

		if len(word) > 0 {
			if lineLen == 0 {
				twc.printAndClear(word)
			} else if lineLen+len(word)+len(spaces) <= maxWidth {
				twc.printAndClear(spaces)
				twc.printAndClear(word)
			} else {
				fmt.Fprintln(twc.W)
				fmt.Fprint(twc.W, prefix)
				twc.printAndClear(word)
			}
		}

		fmt.Fprintln(twc.W)

		firstLinePrefix = strings.Repeat(" ", paraFirstLineIndent)
		maxWidth = paraFirstLineMaxWidth
	}
}

// printAndClear prints a slice of runes as a string and clears the slice
func (twc TWConf) printAndClear(word []rune) []rune {
	if len(word) > 0 {
		fmt.Fprint(twc.W, string(word))
		word = word[:0]
	}
	return word
}
