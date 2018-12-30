package stats

import (
	"strconv"
	"sync"
)

type Total struct {
	mu sync.Mutex

	TotalFiles   int
	TotalLines   int
	CodeLines    int
	CommentLines int
	VoidLines    int
	StatsPerLang map[string]*Total
}

func (ts *Total) Strings() []string {
	return []string{
		strconv.Itoa(ts.TotalFiles),
		strconv.Itoa(ts.TotalLines),
		strconv.Itoa(ts.CodeLines),
		strconv.Itoa(ts.CommentLines),
		strconv.Itoa(ts.VoidLines),
	}
}

func (ts *Total) Append(stats *File) {
	//fmt.Println("appending ", stats, "----- to ", ts)

	ts.TotalFiles++
	ts.TotalLines += stats.TotalLines
	ts.CodeLines += stats.CodeLines
	ts.CommentLines += stats.CommentLines
	ts.VoidLines += stats.VoidLines

	ts.StatsPerLang[stats.Lang.Name].TotalFiles++
	ts.StatsPerLang[stats.Lang.Name].CodeLines += stats.CodeLines
	ts.StatsPerLang[stats.Lang.Name].VoidLines += stats.VoidLines
	ts.StatsPerLang[stats.Lang.Name].TotalLines += stats.TotalLines
	ts.StatsPerLang[stats.Lang.Name].CommentLines += stats.CommentLines

	//fmt.Println("appended.", ts)
}

func (ts *Total) AppendSafe(stats *File) {
	//fmt.Println("appending ", stats, "----- to ", ts)

	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.TotalFiles++
	ts.TotalLines += stats.TotalLines
	ts.CodeLines += stats.CodeLines
	ts.CommentLines += stats.CommentLines
	ts.VoidLines += stats.VoidLines

	ts.StatsPerLang[stats.Lang.Name].TotalFiles++
	ts.StatsPerLang[stats.Lang.Name].CodeLines += stats.CodeLines
	ts.StatsPerLang[stats.Lang.Name].VoidLines += stats.VoidLines
	ts.StatsPerLang[stats.Lang.Name].TotalLines += stats.TotalLines
	ts.StatsPerLang[stats.Lang.Name].CommentLines += stats.CommentLines
	//fmt.Println("appended.", ts)

}
