package twrap_test

import (
	"bytes"
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
	"github.com/nickwells/twrap.mod/twrap"
)

func TestWrap(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		indent  int
		text    string
		expText string
	}{
		{
			ID:   testhelper.MkID("no indent"),
			text: "123 567 901 345 789 123 567 901",
			expText: `123 567 901 345 789
123 567 901
`,
		},
		{
			ID:   testhelper.MkID("no indent - short text (WS at end)"),
			text: "123 567 901 345 789 ",
			expText: `123 567 901 345 789
`,
		},
		{
			ID:   testhelper.MkID("no indent - leading WS"),
			text: "   123 ",
			expText: `   123
`,
		},
		{
			ID:      testhelper.MkID("no indent - empty string"),
			text:    "",
			expText: ``,
		},
		{
			ID:   testhelper.MkID("no indent - multiple WS in middle"),
			text: "123   456",
			expText: `123   456
`,
		},
		{
			ID:     testhelper.MkID("indent = 5"),
			indent: 5,
			text:   "123 567 901 345 789 123 567 901",
			expText: `     123 567 901 345
     789 123 567 901
`,
		},
		{
			ID:     testhelper.MkID("big indent - print min chars"),
			indent: 15,
			text:   "123 567 901 345 789 123 567 901",
			expText: `               123 567
               901 345
               789 123
               567 901
`,
		},
		{
			ID:   testhelper.MkID("long word"),
			text: "aaaaaaaaaaaaaaaaaaaaaa bbb ccc",
			expText: `aaaaaaaaaaaaaaaaaaaaaa
bbb ccc
`,
		},
	}

	for _, tc := range testCases {
		var buf bytes.Buffer

		twc, err := twrap.NewTWConf(
			twrap.TWConfOptSetWriter(&buf),
			twrap.TWConfOptSetMinChars(10),
			twrap.TWConfOptSetTargetLineLen(20))
		if err != nil {
			t.Fatal(tc.IDStr(), ": Couldn't create the TWConf: ", err)
		}

		twc.Wrap(tc.text, tc.indent)
		testhelper.DiffString(t, tc.IDStr(), "wrapped text",
			buf.String(), tc.expText)
	}
}

func TestWrap2Indent(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		firstLineIndent int
		otherLineIndent int
		text            string
		expText         string
	}{
		{
			ID:              testhelper.MkID("indents: 2, 3"),
			firstLineIndent: 2,
			otherLineIndent: 3,
			text:            "aaa bbb ccc ddd eee fff ggg hhh iii jjj",
			expText: `  aaa bbb ccc ddd
   eee fff ggg hhh
   iii jjj
`,
		},
	}

	for _, tc := range testCases {
		var buf bytes.Buffer

		twc, err := twrap.NewTWConf(
			twrap.TWConfOptSetWriter(&buf),
			twrap.TWConfOptSetMinChars(10),
			twrap.TWConfOptSetTargetLineLen(20))
		if err != nil {
			t.Fatal(tc.IDStr(), ": Couldn't create the TWConf: ", err)
		}

		twc.Wrap2Indent(tc.text, tc.firstLineIndent, tc.otherLineIndent)
		testhelper.DiffString(t, tc.IDStr(), "wrapped text",
			buf.String(), tc.expText)
	}
}

func TestWrap3Indent(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		firstLineIndent     int
		paraFirstLineIndent int
		otherLineIndent     int
		text                string
		expText             string
	}{
		{
			ID:                  testhelper.MkID("indents: 2, 3, 4"),
			firstLineIndent:     2,
			paraFirstLineIndent: 3,
			otherLineIndent:     4,
			text:                "aaa bbb ccc\nddd eee fff ggg hhh iii jjj",
			expText: `  aaa bbb ccc
   ddd eee fff ggg
    hhh iii jjj
`,
		},
		{
			ID:                  testhelper.MkID("para with leading '# '"),
			firstLineIndent:     10,
			paraFirstLineIndent: 10,
			otherLineIndent:     10,
			text:                "aaa bbb\n# ddd eee fff ggg hhh iii jjj",
			expText: `          aaa bbb
          # ddd eee
          fff ggg
          hhh iii
          jjj
`,
		},
		{
			ID:                  testhelper.MkID("para with leading '- '"),
			firstLineIndent:     10,
			paraFirstLineIndent: 10,
			otherLineIndent:     10,
			text:                "aaa bbb\n- ddd eee fff ggg hhh iii jjj",
			expText: `          aaa bbb
          - ddd eee
            fff ggg
            hhh iii
            jjj
`,
		},
		{
			ID:                  testhelper.MkID("para with leading '* '"),
			firstLineIndent:     10,
			paraFirstLineIndent: 10,
			otherLineIndent:     10,
			text:                "aaa bbb\n* ddd eee fff ggg hhh iii jjj",
			expText: `          aaa bbb
          * ddd eee
            fff ggg
            hhh iii
            jjj
`,
		},
		{
			ID:                  testhelper.MkID("para with leading '+ '"),
			firstLineIndent:     10,
			paraFirstLineIndent: 10,
			otherLineIndent:     10,
			text:                "aaa bbb\n+ ddd eee fff ggg hhh iii jjj",
			expText: `          aaa bbb
          + ddd eee
            fff ggg
            hhh iii
            jjj
`,
		},
		{
			ID: testhelper.MkID(
				"para with leading '+ ' but different otherLineIndent"),
			firstLineIndent:     10,
			paraFirstLineIndent: 10,
			otherLineIndent:     9,
			text:                "aaa bbb\n+ ddd eee fff ggg hhh iii jjj",
			expText: `          aaa bbb
          + ddd eee
         fff ggg hhh
         iii jjj
`,
		},
	}

	for _, tc := range testCases {
		var buf bytes.Buffer

		twc, err := twrap.NewTWConf(
			twrap.TWConfOptSetWriter(&buf),
			twrap.TWConfOptSetMinChars(10),
			twrap.TWConfOptSetTargetLineLen(20))
		if err != nil {
			t.Fatal(tc.IDStr(), ": Couldn't create the TWConf: ", err)
		}

		twc.Wrap3Indent(tc.text,
			tc.firstLineIndent, tc.paraFirstLineIndent, tc.otherLineIndent)
		testhelper.DiffString(t, tc.IDStr(), "wrapped text",
			buf.String(), tc.expText)
	}
}
