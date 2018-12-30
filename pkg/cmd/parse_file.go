package cmd

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/olekukonko/tablewriter"

	"github.com/a-hilaly/cs/pkg/lang"
	"github.com/a-hilaly/cs/pkg/parser"
)

func ParseFile(out io.Writer, filepath string) (*tablewriter.Table, error) {
	dtc := lang.NewDetector()
	lang, _ := dtc.Detect(filepath)

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading file: %v", bytes)
	}

	s, err := parser.ParseFileBytes(bytes, lang)
	if err != nil {
		return nil, fmt.Errorf("parsing file: %v", err)
	}

	data := [][]string{
		append([]string{s.Lang.Name, "1"}, s.Strings()...),
	}

	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"Language", "Files", "Total", "Code", "Comment", "Blank"})

	for _, v := range data {
		table.Append(v)
	}

	return table, nil
}
