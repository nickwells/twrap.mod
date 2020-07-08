package twrap_test

import (
	"fmt"

	"github.com/nickwells/twrap.mod/twrap"
)

// ExampleTWConf_Wrap provides an exaple of how the twrap.Wrap method might
// be used
func ExampleTWConf_Wrap() {
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

// ExampleTWConf_Wrap2Indent provides an exaple of how the
// twrap.Wrap2Indent method might be used
func ExampleTWConf_Wrap2Indent() {
	twc, err := twrap.NewTWConf()
	if err != nil {
		fmt.Println("cannot construct the TWConf object:", err)
		return
	}
	twc.Wrap2Indent("list item: A description of the item such as"+
		" might appear in a menu.\n"+
		"item2: A description of the second item in the menu. Note the"+
		" indentation of the start of the line and the following text.",
		20, 24)

	// Output:
	//                     list item: A description of the item such as might appear in
	//                         a menu.
	//                     item2: A description of the second item in the menu. Note
	//                         the indentation of the start of the line and the
	//                         following text.
}

// ExampleTWConf_WrapPrefixed provides an exaple of how the
// twrap.WrapPrefixed method might be used
func ExampleTWConf_WrapPrefixed() {
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

// ExampleTWConf_WrapPrefixed_multiPara provides an exaple of how the
// twrap.WrapPrefixed method might be used
func ExampleTWConf_WrapPrefixed_multiPara() {
	twc, err := twrap.NewTWConf()
	if err != nil {
		fmt.Println("cannot construct the TWConf object:", err)
		return
	}
	twc.WrapPrefixed("See also: ",
		"The github.com/nickwells/param.mod/param/phelp package makes"+
			" extensive use of the twrap package"+
			"\n\n"+
			"As do several of the utilities in .../utilities",
		10)

	// Output:
	//           See also: The github.com/nickwells/param.mod/param/phelp package makes
	//                     extensive use of the twrap package
	//
	//                     As do several of the utilities in .../utilities
}
