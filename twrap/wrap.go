package twrap

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

const (
	DfltMinCharsToPrint = 30
	DfltTargetLineLen   = 80
)

// TWConf holds the configuration for a text wrapper
type TWConf struct {
	W               io.Writer
	MinCharsToPrint int
	TargetLineLen   int
}

// TWConfOptFunc is the signature of the function that is passed to the
// NewTWConf function to override the default values
type TWConfOptFunc func(*TWConf) error

// TWConfOptSetWriter - see SetWriter
func TWConfOptSetWriter(w io.Writer) TWConfOptFunc {
	return SetWriter(w)
}

// SetWriter returns a TWConfOptFunc suitable for passing to
// NewTWConf which will set the MinCharsToPrint
func SetWriter(w io.Writer) TWConfOptFunc {
	return func(twc *TWConf) error {
		if w == nil {
			return errors.New("the Writer must not be nil")
		}
		twc.W = w
		return nil
	}
}

// TWConfOptSetMinChars - see SetMinChars
func TWConfOptSetMinChars(n int) TWConfOptFunc {
	return SetMinChars(n)
}

// SetMinChars returns a TWConfOptFunc suitable for passing to
// NewTWConf which will set the MinCharsToPrint
func SetMinChars(n int) TWConfOptFunc {
	return func(twc *TWConf) error {
		if n < 0 {
			return errors.New(
				"the minimum number of chars to print on a line must be >= 0")
		}
		twc.MinCharsToPrint = n
		return nil
	}
}

// TWConfOptSetTargetLineLen - see SetTargetLineLen
func TWConfOptSetTargetLineLen(n int) TWConfOptFunc {
	return SetTargetLineLen(n)
}

// SetTargetLineLen returns a TWConfOptFunc suitable for passing to
// NewTWConf which will set the TargetLineLen
func SetTargetLineLen(n int) TWConfOptFunc {
	return func(twc *TWConf) error {
		if n <= 0 {
			return errors.New("the target line length must be > 0")
		}
		twc.TargetLineLen = n
		return nil
	}
}

// NewTWConf constructs a TWConf with the default values. To override the
// default values pass the appropriate option functions. If any of the
// option funcs returns an error the error is returned and a nil value
func NewTWConf(opts ...TWConfOptFunc) (*TWConf, error) {
	twc := &TWConf{
		W:               os.Stdout,
		MinCharsToPrint: DfltMinCharsToPrint,
		TargetLineLen:   DfltTargetLineLen,
	}

	for _, o := range opts {
		err := o(twc)
		if err != nil {
			return nil, err
		}
	}
	if twc.MinCharsToPrint > twc.TargetLineLen {
		return nil,
			fmt.Errorf("the minimum number of characters to print (%d)"+
				" must be less than the target line length (%d)",
				twc.MinCharsToPrint, twc.TargetLineLen)
	}
	return twc, nil
}

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
func (twc TWConf) Wrap2Indent(text string, firstLineIndent, otherLineIndent int) {
	twc.Wrap3Indent(text, firstLineIndent, firstLineIndent, otherLineIndent)
}

// Wrap3Indent will print the text as with Wrap. The first line printed will
// be indented by the first supplied indent, thereafter the first line of
// each paragraph will be indented by the second supplied indent and any
// other lines will be indented by the third indent value.
func (twc TWConf) Wrap3Indent(text string, firstLineIndent, paraFirstLineIndent, otherLineIndent int) {
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
