package twrap

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
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

// TWConfOptSetWriter returns a TWConfOptFunc suitable for passing to
// TWConfNew which will set the writer
func TWConfOptSetWriter(w io.Writer) TWConfOptFunc {
	return func(twc *TWConf) error {
		if w == nil {
			return errors.New("the Writer must not be nil")
		}
		twc.W = w
		return nil
	}
}

// TWConfOptSetMinChars returns a TWConfOptFunc suitable for passing to
// TWConfNew which will set the MinCharsToPrint
func TWConfOptSetMinChars(n int) TWConfOptFunc {
	return func(twc *TWConf) error {
		if n < 0 {
			return errors.New(
				"the minimum number of chars to print on a line must be >= 0")
		}
		twc.MinCharsToPrint = n
		return nil
	}
}

// TWConfOptSetTargetLineLen returns a TWConfOptFunc suitable for passing to
// TWConfNew which will set the TargetLineLen
func TWConfOptSetTargetLineLen(n int) TWConfOptFunc {
	return func(twc *TWConf) error {
		if n <= 0 {
			return errors.New("the target line length must be > 0")
		}
		twc.TargetLineLen = n
		return nil
	}
}

// NewTWConf constructs a TWConf with the default values. To override the
// default values pass the appropriate opttion functions. If any of the
// option funcs returns an error the error is returned and a nil value
func NewTWConf(opts ...TWConfOptFunc) (*TWConf, error) {
	twc := &TWConf{
		W:               os.Stdout,
		MinCharsToPrint: 30,
		TargetLineLen:   80,
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
				" should be greater than the target line length (%d)",
				twc.MinCharsToPrint, twc.TargetLineLen)
	}
	return twc, nil
}

// Wrap will print the text onto the configured writer but wraps and indents
// the text. It will always print at least MinCharsToPrint chars but will try
// to fit the text into TargetLineLen chars.
func (twc TWConf) Wrap(text string, indent int) {
	twc.Wrap2Indent(text, indent, indent)
}

// Wrap2Indent will print the text onto the configured writer but wraps and
// indents the text. The first line printed will have a different indent from
// the other lines. It will always print at least MinCharsToPrint chars but
// will try to fit the text into TargetLineLen chars.
func (twc TWConf) Wrap2Indent(text string, firstLineIndent, otherLineIndent int) {
	firstLineMaxWidth := int(
		math.Max(
			float64(twc.MinCharsToPrint),
			float64(twc.TargetLineLen-firstLineIndent)))
	otherLineMaxWidth := int(
		math.Max(
			float64(twc.MinCharsToPrint),
			float64(twc.TargetLineLen-otherLineIndent)))

	paras := strings.Split(text, "\n")
	prefix := strings.Repeat(" ", otherLineIndent)

	for _, para := range paras {
		fmt.Fprint(twc.W, strings.Repeat(" ", firstLineIndent))
		maxWidth := firstLineMaxWidth

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

		fmt.Fprintln(twc.W) // nolint: errcheck
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
