package twrap

import "fmt"

// Print calls fmt.Fprint passing the TWConf writer
func (twc TWConf) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(twc.W, a...)
}

// Println calls fmt.Fprintln passing the TWConf writer
func (twc TWConf) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(twc.W, a...)
}

// Printf calls fmt.Fprintf passing the TWConf writer
func (twc TWConf) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(twc.W, format, a...)
}
