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
	line1Indent, paraLine1Indent, line2Indent int,
) {
	if text == "" {
		return
	}

	// work out how much space we have between the indent and the end of line
	line1MaxLen := int(
		math.Max(
			float64(twc.MinCharsToPrint),
			float64(twc.TargetLineLen-line1Indent)))
	paraLine1MaxLen := int(
		math.Max(
			float64(twc.MinCharsToPrint),
			float64(twc.TargetLineLen-paraLine1Indent)))
	line2MaxLen := int(
		math.Max(
			float64(twc.MinCharsToPrint),
			float64(twc.TargetLineLen-line2Indent)))

	paras := strings.Split(text, "\n")
	prefix := strings.Repeat(" ", line2Indent)
	line1Prefix := strings.Repeat(" ", line1Indent)
	maxLen := line1MaxLen

	for _, para := range paras {
		fmt.Fprint(twc.W, line1Prefix)

		lineLen := 0
		word := make([]rune, 0, len(para))
		spaces := make([]rune, 0, maxLen)
		for _, r := range para {
			if r == ' ' {
				if len(word) > 0 {
					lineLen, maxLen = twc.printWord(word, spaces,
						prefix,
						lineLen, maxLen, line2MaxLen)
					word = word[:0]
					spaces = spaces[:0]
				}

				spaces = append(spaces, r)
			} else {
				word = append(word, r)
			}
		}

		if len(word) > 0 {
			twc.printWord(word, spaces, prefix, lineLen, maxLen, line2MaxLen)
		}

		fmt.Fprintln(twc.W)

		line1Prefix = strings.Repeat(" ", paraLine1Indent)
		maxLen = paraLine1MaxLen
	}
}

// printWord prints the word and any leading spaces and returns the new line
// length and the new max length
func (twc TWConf) printWord(word, spaces []rune, prefix string, lineLen, maxLen, nextMaxLen int) (int, int) {
	if lineLen == 0 { // always print the 1st word regardless of its length
		fmt.Fprint(twc.W, string(word))
		return len(word), maxLen
	}

	if lineLen+len(word)+len(spaces) <= maxLen { // word & space fit in the line
		lineLen += len(word) + len(spaces)
		fmt.Fprint(twc.W, string(spaces)+string(word))
		return lineLen, maxLen
	}

	fmt.Fprintln(twc.W)
	fmt.Fprint(twc.W, prefix+string(word))
	return len(word), nextMaxLen
}
