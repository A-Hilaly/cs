package parser

import (
	"bytes"

	"github.com/a-hilaly/cs/pkg/lang"
	"github.com/a-hilaly/cs/pkg/stats"
)

var (
	EOL   = []byte("\n")
	SPACE = []byte(" ")
)

type Parser func([]byte, *lang.Spec) (*stats.File, error)

func ParseFileBytes(b []byte, ls *lang.Spec) (*stats.File, error) {
	if ls.Name == lang.UnknownLang.Name {
		return countFileLines(b, ls)
	}

	if ls.Comment.BlockStart == nil && ls.Comment.BlockEnd == nil {
		return parseOneCommentFileBytes(b, ls)
	}

	incomment := false
	totalLines := 0
	emptyLines := 0
	codeLinesCount := 0
	commentLinesCount := 0
	loopSeenBlockStart := false
	blines := bytes.Split(b, EOL)

	for _, bline := range blines {

		totalLines++

		if len(bline) == 0 ||
			len(bytes.Split(bline[0:len(bline)], SPACE)) == len(bline)+1 {

			emptyLines++
			if incomment {

				commentLinesCount++
			} else {

				codeLinesCount++
			}
			continue
		}

		// Check comment block start
		if c := bytes.Index(bline, ls.Comment.BlockStart); c == 0 {

			commentLinesCount++
			incomment = true
			loopSeenBlockStart = true

		} else if c > 1 {

			commentLinesCount++
			codeLinesCount++
			incomment = true
			loopSeenBlockStart = true
		}

		// checking for block end
		// if already detected block comment start should be
		if incomment {

			if c := bytes.Index(bline, ls.Comment.BlockEnd); c == -1 {

				commentLinesCount++

				continue

			} else {

				if !loopSeenBlockStart {
					commentLinesCount++
				}
				incomment = false
				continue
			}
		}

		// Check comment line
		if c := bytes.Index(bline, ls.Comment.One); c == -1 {

			codeLinesCount++
			continue
		} else if c == 0 {

			commentLinesCount++
			continue

		} else {

			if len(bytes.Split(bline[0:c], SPACE)) == c+1 {
				commentLinesCount++
			} else {
				commentLinesCount++
				codeLinesCount++
			}
		}

		if loopSeenBlockStart {
			loopSeenBlockStart = false
		}
	}

	return &stats.File{
		Lang:         ls,
		TotalLines:   totalLines,
		CodeLines:    codeLinesCount,
		CommentLines: commentLinesCount,
		VoidLines:    emptyLines,
	}, nil
}

func countFileLines(b []byte, ls *lang.Spec) (*stats.File, error) {
	return &stats.File{
		Lang:       ls,
		TotalLines: bytes.Count(b, EOL),
	}, nil
}

func parseOneCommentFileBytes(b []byte, ls *lang.Spec) (*stats.File, error) {
	totalLines := 0
	emptyLines := 0
	codeLinesCount := 0
	commentLinesCount := 0
	blines := bytes.Split(b, EOL)

	for _, bline := range blines {

		totalLines++
		if len(bline) == 0 || len(bytes.Split(bline[0:len(bline)], SPACE)) == len(bline)+1 {
			emptyLines++
			codeLinesCount++
			continue
		}

		// Check comment line
		if c := bytes.Index(bline, ls.Comment.One); c == -1 {
			codeLinesCount++
			continue
		} else if c == 0 {
			commentLinesCount++
			continue

		} else {
			if len(bytes.Split(bline[0:c], SPACE)) == c+1 {
				commentLinesCount++
			} else {
				codeLinesCount++
				commentLinesCount++
			}
		}
	}

	return &stats.File{
		Lang:         ls,
		TotalLines:   totalLines,
		VoidLines:    emptyLines,
		CodeLines:    codeLinesCount,
		CommentLines: commentLinesCount,
	}, nil
}
