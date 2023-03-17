package twrap

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

// we break the string to be wrapped into paragraphs on either newlines or
// form feeds
var paraBreakRE = regexp.MustCompile("[\n\f]")

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

// calcMaxLen returns the max line length given the specified indent.
func (twc TWConf) calcMaxLen(indent int) int {
	return int(
		math.Max(
			float64(twc.MinCharsToPrint),
			float64(twc.TargetLineLen-indent)))
}

// isAListItem returns true if the paragraph looks like an item in a
// bulletted list - if the paragraph starts with one of '-', '*' or '+'
// followed by a space
func isAListItem(para string) bool {
	return (strings.HasPrefix(para, "- ") ||
		strings.HasPrefix(para, "* ") ||
		strings.HasPrefix(para, "+ "))
}

// isABreakableSpace returns true if the rune is a space and is not equal to
// the non-breakable space rune
func isABreakableSpace(r rune) bool {
	return unicode.IsSpace(r) && r != '\u00a0' // break on space, not NBSP
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
	var (
		line1MaxLen     = twc.calcMaxLen(line1Indent)
		paraLine1MaxLen = twc.calcMaxLen(paraLine1Indent)
		line2MaxLen     = twc.calcMaxLen(line2Indent)
	)

	paras := paraBreakRE.Split(text, -1)
	line1Prefix := strings.Repeat(" ", line1Indent)
	line2Prefix := strings.Repeat(" ", line2Indent)
	maxLen := line1MaxLen

	for _, para := range paras {
		if para != "" {
			twc.Print(line1Prefix)
		}

		prefix := line2Prefix
		secondLineMaxLen := line2MaxLen
		if isAListItem(para) && line2MaxLen == maxLen {
			prefix += "  "
			secondLineMaxLen = twc.calcMaxLen(line2Indent + 2)
		}

		lineLen := 0
		word := make([]rune, 0, len(para))
		spaces := make([]rune, 0, maxLen)
		for _, r := range para {
			if isABreakableSpace(r) {
				if len(word) > 0 {
					lineLen, maxLen = twc.printWord(word, spaces,
						prefix,
						lineLen, maxLen, secondLineMaxLen)
					word = word[:0]
					spaces = spaces[:0]
				}

				spaces = append(spaces, r)
			} else {
				word = append(word, r)
			}
		}

		if len(word) > 0 {
			twc.printWord(word, spaces,
				prefix,
				lineLen, maxLen, secondLineMaxLen)
		}

		twc.Println()

		line1Prefix = strings.Repeat(" ", paraLine1Indent)
		maxLen = paraLine1MaxLen
	}
}

// printWord prints the word and any leading spaces and returns the new line
// length and the new max length
func (twc TWConf) printWord(word, spaces []rune, prefix string,
	lineLen, maxLen, nextMaxLen int,
) (int, int) {
	if lineLen == 0 {
		// always print 1st word regardless of length (with leading spaces)
		twc.Print(string(spaces) + string(word))
		return len(spaces) + len(word), maxLen
	}

	if lineLen+len(word)+len(spaces) <= maxLen { // word & space fit in the line
		lineLen += len(word) + len(spaces)
		twc.Print(string(spaces) + string(word))
		return lineLen, maxLen
	}

	twc.Println()
	twc.Print(prefix + string(word))
	return len(word), nextMaxLen
}
