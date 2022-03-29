package twrap_test

import (
	"bytes"
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
	"github.com/nickwells/twrap.mod/twrap"
)

const (
	testDataDir = "testdata"
	listSubDir  = "list"
)

var gfcList = testhelper.GoldenFileCfg{
	DirNames:               []string{testDataDir, listSubDir},
	Sfx:                    "txt",
	UpdFlagName:            "upd-list-files",
	KeepBadResultsFlagName: "keep-bad-list-results",
}

func init() {
	gfcList.AddUpdateFlag()
	gfcList.AddKeepBadResultsFlag()
}

func TestList(t *testing.T) {
	list := []string{
		"entry with no path parts",
		"entry1",
		"part1/part2/entry2",
		"part1/part2/entry3",
		"part1/part2/entry4: Very long text that is expected to wrap." +
			" Blah, blah, blah, blah, blah, blah, blah, blah, blah, blah," +
			" blah, blah, blah, blah, blah, blah, blah, blah, blah, blah," +
			" blah, blah, blah, blah, blah, blah, blah, blah, blah, blah," +
			" blah, blah, blah.",
		"part1/part2/entry5",
		"part1/part3/entry6 with some trailing text",
		"part1/part4/entry7",
		"part5/entry8",
		"part5/part6/entry9",
		"/part5/part7/entry10",
		"/part5//part7/entry11",
		"/part5/part7/entry12",
		"/part5/part7/entry13",
		"/part5/part7/entry14",
		"/part5/part7/entry15",
	}
	testCases := []struct {
		testhelper.ID
		list   []string
		indent int
	}{
		{
			ID:   testhelper.MkID("1-entry-indent0"),
			list: list[0:1],
		},
		{
			ID:     testhelper.MkID("1-entry-indent5"),
			list:   list[0:1],
			indent: 5,
		},
		{
			ID:   testhelper.MkID("9-entry-indent0"),
			list: list[0:9],
		},
		{
			ID:     testhelper.MkID("9-entry-indent5"),
			list:   list[0:9],
			indent: 5,
		},
		{
			ID:   testhelper.MkID("10-entry-indent0"),
			list: list[0:10],
		},
		{
			ID:     testhelper.MkID("10-entry-indent5"),
			list:   list[0:10],
			indent: 5,
		},
		{
			ID:   testhelper.MkID("leading-sep-indent0"),
			list: list[10:],
		},
		{
			ID:     testhelper.MkID("leading-sep-indent5"),
			list:   list[10:],
			indent: 5,
		},
	}

	for _, tc := range testCases {
		for _, l := range []struct {
			f    func(twrap.TWConf, []string, int)
			name string
		}{
			{f: twrap.TWConf.List, name: "-List"},
			{f: twrap.TWConf.IdxList, name: "-IdxList"},
			{f: twrap.TWConf.NoRptList, name: "-NoRptList"},
			{f: twrap.TWConf.NoRptPathList, name: "-NoRptPathList"},
			{f: twrap.TWConf.IdxNoRptList, name: "-IdxNoRptList"},
			{f: twrap.TWConf.IdxNoRptPathList, name: "-IdxNoRptPathList"},
		} {
			var b bytes.Buffer
			twc := twrap.NewTWConfOrPanic(twrap.SetWriter(&b))
			l.f(*twc, tc.list, tc.indent)
			gfcList.Check(t, tc.IDStr()+l.name, tc.ID.Name+l.name, b.Bytes())
		}
	}
}
