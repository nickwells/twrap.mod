package twrap

import (
	"fmt"
	"math"
	"os"
	"path"
	"strings"
)

// calcNumDigits calculates the number of digits that will be needed to display the given value
func calcNumDigits(i int) int {
	return int(math.Ceil(math.Log10(float64(i + 1))))
}

// List will print the list of strings, one per line, with the appropriate
// indent and with each item prefixed with the list prefix
func (twc TWConf) List(list []string, indent int) {
	for _, li := range list {
		twc.WrapPrefixed(twc.ListPrefix+" ", li, indent)
	}
}

// IdxList will print the list of strings, one per line, with the
// appropriate indent and with each item prefixed with an index number
func (twc TWConf) IdxList(list []string, indent int) {
	digits := calcNumDigits(len(list))
	for i, li := range list {
		twc.WrapPrefixed(
			fmt.Sprintf("%s %*d: ", twc.ListPrefix, digits, i+1),
			li,
			indent)
	}
}

// NoRptList will print a list of strings with no wrapping. Each list item
// will have any characters from the start of the string which are common to
// the preceding list item replaced with spaces
func (twc TWConf) NoRptList(list []string, indent int) {
	prev := []rune{}
	for _, li := range list {
		twc.Print(strings.Repeat(" ", indent) + twc.ListPrefix + " ")
		prev = twc.printUniqueStrParts(li, prev)
	}
}

// IdxNoRptList will print a list of strings with no wrapping. Each
// list item will be prefixed with an index number and will have any
// characters from the start of the string which are common to the preceding
// list item replaced with spaces
func (twc TWConf) IdxNoRptList(list []string, indent int) {
	digits := calcNumDigits(len(list))
	prev := []rune{}
	for i, li := range list {
		twc.Print(strings.Repeat(" ", indent))
		twc.Print(idxListPrefix(twc.ListPrefix, i+1, digits))
		prev = twc.printUniqueStrParts(li, prev)
	}
}

// NoRptPathList will print a list of strings with no wrapping. Each list
// item is assumed to be a pathname and will have any part of the path
// (except the last) which is the same as the corresponding part of the
// previous list item replaced with spaces. As soon as any part of the path
// differs, the remainder is printed as is.
func (twc TWConf) NoRptPathList(list []string, indent int) {
	prev := []string{}
	for _, li := range list {
		twc.Print(strings.Repeat(" ", indent) + twc.ListPrefix + " ")
		dir, file := path.Split(li)
		prev = twc.printUniqueDirParts(dir, prev)
		twc.Print(file + "\n")
	}
}

// IdxNoRptPathList will print a list of strings with no wrapping. Each list
// item will be prefixed with an index number. Each list item is assumed to
// be a pathname and will have any part of the path (except the last) which
// is the same as the corresponding part of the previous list item replaced
// with spaces. As soon as any part of the path differs, the remainder is
// printed as is.
func (twc TWConf) IdxNoRptPathList(list []string, indent int) {
	digits := calcNumDigits(len(list))
	prev := []string{}
	for i, li := range list {
		twc.Print(strings.Repeat(" ", indent))
		twc.Print(idxListPrefix(twc.ListPrefix, i+1, digits))
		dir, file := path.Split(li)
		prev = twc.printUniqueDirParts(dir, prev)
		twc.Print(file + "\n")
	}
}

// idxListPrefix will return a suitable indexed list prefix
func idxListPrefix(listPfx string, i, digits int) string {
	return fmt.Sprintf("%s %*d: ", listPfx, digits, i)
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
		pathSep := string(os.PathSeparator)
		dirParts = strings.Split(dir, pathSep)
		for i, dp := range dirParts {
			if dp == "" {
				break
			}
			if i >= len(prevParts) || dp != prevParts[i] {
				twc.Print(strings.Join(dirParts[i:], pathSep))
				break
			}
			twc.Print(strings.Repeat(" ", len(dp)+len(pathSep)))
		}
	}
	return dirParts
}
