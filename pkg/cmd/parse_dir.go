package cmd

import (
	"fmt"
	"io"
	"sort"

	"github.com/olekukonko/tablewriter"

	"github.com/a-hilaly/cs/pkg/walker"
)

func WalkAndParseDir(out io.Writer, filepath string) (*tablewriter.Table, error) {
	w := walker.New(&walker.Config{})

	s, err := w.Walk(filepath)
	if err != nil {
		return nil, fmt.Errorf("parsing file: %v", err)
	}

	var keys []string
	for k := range s.StatsPerLang {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	data := [][]string{}
	for _, lang := range keys {
		data = append(data,
			append([]string{lang}, s.StatsPerLang[lang].Strings()...),
		)
	}

	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"Language", "Files", "Total", "Code", "Comment", "Blank"})

	for _, v := range data {
		table.Append(v)
	}

	//table.SetBorder(false)
	table.SetFooter(append([]string{"total"}, s.Strings()...))
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)

	return table, nil
}
