package stats

import (
	"strconv"

	"github.com/a-hilaly/cs/pkg/lang"
)

type File struct {
	Lang         *lang.Spec
	TotalLines   int
	CodeLines    int
	CommentLines int
	VoidLines    int
}

func (fs *File) Strings() []string {
	return []string{
		strconv.Itoa(fs.TotalLines),
		strconv.Itoa(fs.CodeLines),
		strconv.Itoa(fs.CommentLines),
		strconv.Itoa(fs.VoidLines),
	}
}
