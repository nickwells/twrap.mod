package twrap_test

import (
	"fmt"

	"github.com/nickwells/twrap.mod/twrap"
)

// ExampleWrap provides an exaple of how the twrap.Wrap method might be used
func ExampleWrap() {
	twc, err := twrap.NewTWConf()
	if err != nil {
		fmt.Println("cannot construct the TWConf object:", err)
		return
	}
	twc.Wrap("\nThe quality of mercy is not strained."+
		" It droppeth as the gentle rain from heaven."+
		" Upon the place beneath.", 20)

	// Output:
	//                     The quality of mercy is not strained. It droppeth as the
	//                     gentle rain from heaven. Upon the place beneath.
}

// ExampleWrapPrefixed provides an exaple of how the twrap.WrapPrefixed method
// might be used
func ExampleWrapPrefixed() {
	twc, err := twrap.NewTWConf()
	if err != nil {
		fmt.Println("cannot construct the TWConf object:", err)
		return
	}
	twc.WrapPrefixed("Portia: ", "The quality of mercy is not strained."+
		" It droppeth as the gentle rain from heaven."+
		" Upon the place beneath.", 20)

	// Output:
	//                     Portia: The quality of mercy is not strained. It droppeth as
	//                             the gentle rain from heaven. Upon the place beneath.
}
