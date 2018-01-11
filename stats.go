package main

import "fmt"

type Stats struct {
    TotalLines        int
    TotalChars        int
    CodeLines         int
    CommentLines      int
    CodeCharsCount    int
    CommentCharsCount int
    VoidLines         int
}

func (s *Stats) AppendFromStats(stats Stats) {
    s.TotalLines += stats.TotalLines
    s.TotalChars += stats.TotalChars
    s.CodeLines += stats.CodeLines
    s.CommentLines += stats.CommentLines
    s.CodeCharsCount += stats.CodeCharsCount
    s.CommentCharsCount += stats.CommentCharsCount
    s.VoidLines += stats.VoidLines
}

func (s *Stats) Show() {
    fmt.Println("|*Lines")
    fmt.Println("|   Total lines      :", s.TotalLines)
    fmt.Println("|   Code lines       :", s.CodeLines)
    fmt.Println("|   Comment lines    :", s.CommentLines)
    fmt.Println("|   Void lines       :", s.VoidLines)
    fmt.Println("|   Unknown lines    :", s.TotalLines - s.CodeLines -s.CommentLines - s.VoidLines)
    fmt.Println("|*Chars")
    fmt.Println("|   Total Chars      :", s.TotalChars)
    fmt.Println("|   Code Chars       :", s.CodeCharsCount)
    fmt.Println("|   Comment Chars    :", s.CommentCharsCount)
    fmt.Println("|   Unknown Chars    :", s.TotalChars - s.CommentCharsCount - s.CodeCharsCount)

}
