package twrap

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/nickwells/mathutil.mod/v2/mathutil"
)

// List will print the list of strings, one per line, with the appropriate
// indent and with each item prefixed with the list prefix
func (twc TWConf) List(list []string, indent int) {
	for _, li := range list {
		twc.WrapPrefixed(twc.ListPrefix, li, indent)
	}
}

// ListItem calls List it is simply a more convenient interface
func (twc TWConf) ListItem(indent int, list ...string) {
	twc.List(list, indent)
}

// IdxList will print the list of strings, one per line, with the
// appropriate indent and with each item prefixed with an index number
func (twc TWConf) IdxList(list []string, indent int) {
	digits := mathutil.Digits(len(list))
	for i, li := range list {
		twc.WrapPrefixed(
			idxListPrefix(twc.ListPrefix, i+1, digits),
			li,
			indent)
	}
}

// IdxListItem calls IdxList it is simply a more convenient interface
func (twc TWConf) IdxListItem(indent int, list ...string) {
	twc.IdxList(list, indent)
}

// NoRptList will print a list of strings with no wrapping. Each list item
// will have any characters from the start of the string which are common to
// the preceding list item replaced with spaces
func (twc TWConf) NoRptList(list []string, indent int) {
	prev := []rune{}

	for _, li := range list {
		twc.Print(strings.Repeat(" ", indent) + twc.ListPrefix)
		prev = twc.printUniqueStrParts(li, prev)
	}
}

// NoRptListItem calls NoRptList it is simply a more convenient interface
func (twc TWConf) NoRptListItem(indent int, list ...string) {
	twc.NoRptList(list, indent)
}

// IdxNoRptList will print a list of strings with no wrapping. Each
// list item will be prefixed with an index number and will have any
// characters from the start of the string which are common to the preceding
// list item replaced with spaces
func (twc TWConf) IdxNoRptList(list []string, indent int) {
	digits := mathutil.Digits(len(list))
	prev := []rune{}

	for i, li := range list {
		twc.Print(strings.Repeat(" ", indent))
		twc.Print(idxListPrefix(twc.ListPrefix, i+1, digits))
		prev = twc.printUniqueStrParts(li, prev)
	}
}

// IdxNoRptListItem calls IdxNoRptList it is simply a more convenient
// interface
func (twc TWConf) IdxNoRptListItem(indent int, list ...string) {
	twc.IdxNoRptList(list, indent)
}

// NoRptPathList will print a list of strings with no wrapping. Each list
// item is assumed to be a pathname and will have any part of the path
// (except the last) which is the same as the corresponding part of the
// previous list item replaced with spaces. As soon as any part of the path
// differs, the remainder is printed as is.
func (twc TWConf) NoRptPathList(list []string, indent int) {
	prev := []string{}

	for _, li := range list {
		twc.Print(strings.Repeat(" ", indent) + twc.ListPrefix)

		dir, file := filepath.Split(li)
		prev = twc.printUniqueDirParts(dir, prev)

		twc.Print(file + "\n")
	}
}

// NoRptPathListItem calls NoRptPathList it is simply a more convenient
// interface
func (twc TWConf) NoRptPathListItem(indent int, list ...string) {
	twc.NoRptPathList(list, indent)
}

// IdxNoRptPathList will print a list of strings with no wrapping. Each list
// item will be prefixed with an index number. Each list item is assumed to
// be a pathname and will have any part of the path (except the last) which
// is the same as the corresponding part of the previous list item replaced
// with spaces. As soon as any part of the path differs, the remainder is
// printed as is.
func (twc TWConf) IdxNoRptPathList(list []string, indent int) {
	digits := mathutil.Digits(len(list))
	prev := []string{}

	for i, li := range list {
		twc.Print(strings.Repeat(" ", indent))
		twc.Print(idxListPrefix(twc.ListPrefix, i+1, digits))

		dir, file := filepath.Split(li)
		prev = twc.printUniqueDirParts(dir, prev)

		twc.Print(file + "\n")
	}
}

// IdxNoRptPathListItem calls IdxNoRptPathList it is simply a more convenient
// interface
func (twc TWConf) IdxNoRptPathListItem(indent int, list ...string) {
	twc.IdxNoRptPathList(list, indent)
}

// idxListPrefix will return a suitable indexed list prefix
func idxListPrefix(listPfx string, i, digits int) string {
	return fmt.Sprintf("%s%*d: ", listPfx, digits, i)
}

// printUniqueStrParts replaces the leading part of the string that is the
// same as in prev with the equivalent number of spaces. It returns the slice
// of runes for setting prev for the next call
func (twc TWConf) printUniqueStrParts(s string, prev []rune) []rune {
	ra := []rune(s)
	out := make([]rune, 0, len(s))

	for i, r := range ra {
		if i >= len(prev) {
			out = append(out, ra[i:]...)
			break
		}

		if r != prev[i] {
			out = append(out, ra[i:]...)
			break
		}

		out = append(out, ' ')
	}

	twc.Print(string(out) + "\n")

	return ra
}

// printUniqueDirParts replaces those leading parts of the directory that are
// the same as in the prevParts with the equivalent number of spaces. It
// returns the split directory for setting the prevParts for the next call
func (twc TWConf) printUniqueDirParts(dir string, prevParts []string) []string {
	dirParts := []string{}

	if dir != "" {
		dir = filepath.Clean(dir)
		pathSep := string(filepath.Separator)

		dirParts = strings.Split(dir, pathSep)
		for i, dp := range dirParts {
			if i >= len(prevParts) || dp != prevParts[i] {
				twc.Print(strings.Join(dirParts[i:], pathSep))
				twc.Print(pathSep)

				break
			}

			twc.Print(strings.Repeat(" ", len(dp)+len(pathSep)))
		}
	}

	return dirParts
}
