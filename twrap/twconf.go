package twrap

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	DfltMinCharsToPrint = 30
	DfltTargetLineLen   = 80
	DfltListPrefix      = "- "
)

// TWConf holds the configuration for a text wrapper
type TWConf struct {
	W               io.Writer
	MinCharsToPrint int
	TargetLineLen   int
	ListPrefix      string
}

// TWConfOptFunc is the signature of the function that is passed to the
// NewTWConf function to override the default values
type TWConfOptFunc func(*TWConf) error

// TWConfOptSetWriter
//
// Deprecated: use SetWriter instead
func TWConfOptSetWriter(w io.Writer) TWConfOptFunc {
	return SetWriter(w)
}

// SetWriter returns a TWConfOptFunc suitable for passing to NewTWConf which
// will set the Writer from the default value of os.Stdout. The writer must
// not be nil.
func SetWriter(w io.Writer) TWConfOptFunc {
	return func(twc *TWConf) error {
		if w == nil {
			return errors.New("the Writer must not be nil")
		}

		twc.W = w

		return nil
	}
}

// TWConfOptSetMinChars
//
// Deprecated: use SetMinChars instead
func TWConfOptSetMinChars(n int) TWConfOptFunc {
	return SetMinChars(n)
}

// SetMinChars returns a TWConfOptFunc suitable for passing to NewTWConf
// which will set the MinCharsToPrint. The value must be greater or equal to
// zero.
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

// SetListPrefix returns a TWConfOptFunc suitable for passing to
// NewTWConf which will set the ListPrefix
func SetListPrefix(pfx string) TWConfOptFunc {
	return func(twc *TWConf) error {
		twc.ListPrefix = pfx
		return nil
	}
}

// TWConfOptSetTargetLineLen
//
// Deprecated: use SetTargetLineLen instead
func TWConfOptSetTargetLineLen(n int) TWConfOptFunc {
	return SetTargetLineLen(n)
}

// SetTargetLineLen returns a TWConfOptFunc suitable for passing to NewTWConf
// which will set the TargetLineLen. The new target line length must be
// greater than zero.
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
		ListPrefix:      DfltListPrefix,
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
				" must not be greater than the target line length (%d)",
				twc.MinCharsToPrint, twc.TargetLineLen)
	}

	return twc, nil
}

// NewTWConfOrPanic constructs a TWConf using the NewTWConf func but will
// panic if any error is returned. This is more convenient when constructing
// a TWConf as it's a programming error if bad parameters have been passed
// and so a panic is appropriate.
func NewTWConfOrPanic(opts ...TWConfOptFunc) *TWConf {
	twc, err := NewTWConf(opts...)
	if err != nil {
		panic(fmt.Errorf("Couldn't create a new TWConf: %w", err))
	}

	return twc
}
