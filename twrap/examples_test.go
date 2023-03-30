package twrap_test

import (
	"github.com/nickwells/twrap.mod/twrap"
)

// ExampleTWConf_Wrap provides an exaple of how the twrap.Wrap method might
// be used.
//
// Note that it uses the NewTWConfOrPanic func. This will panic if there are
// any problems with the optional parameters passed to it. If you want to
// handle any errors yourself then use the NewTWConf func instead. This will
// return a non-nil error if there are any problems with the parameters and
// the returned TWConf value shoud not be used.
func ExampleTWConf_Wrap() {
	twc := twrap.NewTWConfOrPanic()

	twc.Wrap("\nThe quality of mercy is not strained."+
		" It droppeth as the gentle rain from heaven."+
		" Upon the place beneath.", 20)

	// Output:
	//                     The quality of mercy is not strained. It droppeth as the
	//                     gentle rain from heaven. Upon the place beneath.
}

// ExampleTWConf_Wrap_shortLines provides an exaple of how the twrap.Wrap
// method might be used. The TWConf is created with a shorter line length.
//
// Note that it uses the NewTWConfOrPanic func. This will panic if there are
// any problems with the optional parameters passed to it. If you want to
// handle any errors yourself then use the NewTWConf func instead. This will
// return a non-nil error if there are any problems with the parameters and
// the returned TWConf value shoud not be used.
func ExampleTWConf_Wrap_shortLines() {
	twc := twrap.NewTWConfOrPanic(twrap.SetTargetLineLen(60))

	twc.Wrap("\nThe quality of mercy is not strained."+
		" It droppeth as the gentle rain from heaven."+
		" Upon the place beneath.", 20)

	// Output:
	//                     The quality of mercy is not strained. It
	//                     droppeth as the gentle rain from heaven.
	//                     Upon the place beneath.
}

// ExampleTWConf_Wrap2Indent provides an exaple of how the
// twrap.Wrap2Indent method might be used.
//
// Note that it uses the NewTWConfOrPanic func. This will panic if there are
// any problems with the optional parameters passed to it. If you want to
// handle any errors yourself then use the NewTWConf func instead. This will
// return a non-nil error if there are any problems with the parameters and
// the returned TWConf value shoud not be used.
func ExampleTWConf_Wrap2Indent() {
	twc := twrap.NewTWConfOrPanic()

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
// twrap.WrapPrefixed method might be used.
//
// Note that it uses the NewTWConfOrPanic func. This will panic if there are
// any problems with the optional parameters passed to it. If you want to
// handle any errors yourself then use the NewTWConf func instead. This will
// return a non-nil error if there are any problems with the parameters and
// the returned TWConf value shoud not be used.
func ExampleTWConf_WrapPrefixed() {
	twc := twrap.NewTWConfOrPanic()

	twc.WrapPrefixed("Portia: ", "The quality of mercy is not strained."+
		" It droppeth as the gentle rain from heaven."+
		" Upon the place beneath.", 20)

	// Output:
	//                     Portia: The quality of mercy is not strained. It droppeth as
	//                             the gentle rain from heaven. Upon the place beneath.
}

// ExampleTWConf_WrapPrefixed_multiPara provides an exaple of how the
// twrap.WrapPrefixed method might be used.
//
// Note that it uses the NewTWConfOrPanic func. This will panic if there are
// any problems with the optional parameters passed to it. If you want to
// handle any errors yourself then use the NewTWConf func instead. This will
// return a non-nil error if there are any problems with the parameters and
// the returned TWConf value shoud not be used.
func ExampleTWConf_WrapPrefixed_multiPara() {
	twc := twrap.NewTWConfOrPanic()

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

// ExampleTWConf_ListItem provides an exaple of how the twrap.ListItem method
// might be used.
//
// Note that there is an alternative twrap.List method with the same
// behaviour, ListItem simply provides an alternative interface to the same
// behaviour.
//
// Note that it uses the NewTWConfOrPanic func. This will panic if there are
// any problems with the optional parameters passed to it. If you want to
// handle any errors yourself then use the NewTWConf func instead. This will
// return a non-nil error if there are any problems with the parameters and
// the returned TWConf value shoud not be used.
func ExampleTWConf_ListItem() {
	twc := twrap.NewTWConfOrPanic()

	twc.ListItem(4, "The quality of mercy is not strained.",
		"It droppeth as the gentle rain from heaven.",
		"Upon the place beneath.")

	// Output:
	//     - The quality of mercy is not strained.
	//     - It droppeth as the gentle rain from heaven.
	//     - Upon the place beneath.
}

// ExampleTWConf_IdxListItem provides an exaple of how the twrap.IdxListItem
// method might be used.
//
// Note that there is an alternative twrap.IdxList method with the same
// behaviour, IdxListItem simply provides an alternative interface to the
// same behaviour.
//
// Note that it uses the NewTWConfOrPanic func. This will panic if there are
// any problems with the optional parameters passed to it. If you want to
// handle any errors yourself then use the NewTWConf func instead. This will
// return a non-nil error if there are any problems with the parameters and
// the returned TWConf value shoud not be used.
func ExampleTWConf_IdxListItem() {
	twc := twrap.NewTWConfOrPanic()

	twc.IdxListItem(4, "The quality of mercy is not strained.",
		"It droppeth as the gentle rain from heaven.",
		"Upon the place beneath.")

	// Output:
	//     - 1: The quality of mercy is not strained.
	//     - 2: It droppeth as the gentle rain from heaven.
	//     - 3: Upon the place beneath.
}

// ExampleTWConf_IdxListItem_manyItems provides an exaple of how the
// twrap.IdxListItem method might be used. It highlights how the space for
// the digits automatically adjusts to the length of the list.
//
// Note that there is an alternative twrap.IdxList method with the same
// behaviour, IdxListItem simply provides an alternative interface to the
// same behaviour.
//
// Note that it uses the NewTWConfOrPanic func. This will panic if there are
// any problems with the optional parameters passed to it. If you want to
// handle any errors yourself then use the NewTWConf func instead. This will
// return a non-nil error if there are any problems with the parameters and
// the returned TWConf value shoud not be used.
func ExampleTWConf_IdxListItem_manyItems() {
	twc := twrap.NewTWConfOrPanic()

	twc.IdxListItem(4, "first item",
		"second item",
		"third item",
		"fourth",
		"and another",
		"and yet another",
		"seven",
		"eight",
		"nine",
		"ten",
		"that's sufficient")

	// Output:
	//     -  1: first item
	//     -  2: second item
	//     -  3: third item
	//     -  4: fourth
	//     -  5: and another
	//     -  6: and yet another
	//     -  7: seven
	//     -  8: eight
	//     -  9: nine
	//     - 10: ten
	//     - 11: that's sufficient
}

// ExampleTWConf_NoRptListItem provides an exaple of how the
// twrap.NoRptListItem method might be used.
//
// Note that there is an alternative twrap.NoRptList method with the same
// behaviour, NoRptListItem simply provides an alternative interface to the
// same behaviour.
//
// Note that it uses the NewTWConfOrPanic func. This will panic if there are
// any problems with the optional parameters passed to it. If you want to
// handle any errors yourself then use the NewTWConf func instead. This will
// return a non-nil error if there are any problems with the parameters and
// the returned TWConf value shoud not be used.
func ExampleTWConf_NoRptListItem() {
	twc := twrap.NewTWConfOrPanic()

	twc.NoRptListItem(4,
		"/abc/def/123",
		"/abc/def/124",
		"/abc/ghi/123")

	// Output:
	//     - /abc/def/123
	//     -            4
	//     -      ghi/123
}

// ExampleTWConf_NoRptPathListItem provides an exaple of how the
// twrap.NoRptPathListItem method might be used.
//
// Note that there is an alternative twrap.NoRptPathList method with the same
// behaviour, NoRptPathListItem simply provides an alternative interface to
// the same behaviour.
//
// Note that it uses the NewTWConfOrPanic func. This will panic if there are
// any problems with the optional parameters passed to it. If you want to
// handle any errors yourself then use the NewTWConf func instead. This will
// return a non-nil error if there are any problems with the parameters and
// the returned TWConf value shoud not be used.
func ExampleTWConf_NoRptPathListItem() {
	twc := twrap.NewTWConfOrPanic()

	twc.NoRptPathListItem(4,
		"/abc/def/123",
		"/abc/def/124",
		"/abc/ghi/123")

	// Output:
	//     - /abc/def/123
	//     -          124
	//     -      ghi/123
}
