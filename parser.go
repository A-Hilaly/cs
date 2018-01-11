package main

import (
    "os"
    "bufio"
    "strings"
    "time"
)

type FileParser struct {
    State           bool
    Size            int64
    ModTime         time.Time
    Lang            Language
    Content         [][]byte
    TotalLineNumber int
    TotalCharNumber int
    Index           int
    CIndex          int
    CLine           int
    ParsingCode     bool
    ParsingComment  bool
    CodeLenght      int
    CommentLenght   int
    CodeLines       int
    CommentLines    int
    VoidLines       int
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func OpenFile(path string) (*os.File, error) {
    file, err := os.Open(path)
    if err != nil {
        //log.Fatal(err)
        return nil, Error("Open error")
    }
    return file, nil
}

func (fp *FileParser) LoadStats(file *os.File) error {
    stats, err := file.Stat()
    if err != nil {
        return Error("Stats Error")
    }
    fp.Size = stats.Size()
    fp.ModTime = stats.ModTime()
    return nil
}

func (fp *FileParser) LoadContent(filedir string) error {
    file, err := OpenFile(filedir)
    defer file.Close()
    if err != nil {
        //log.Fatal(err)
        return Error("Open error")
    }
    err = fp.LoadStats(file)
    if err != nil {
        return Error("Stats Error")
    }
    content := make([][]byte, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        content = append(content, []byte(line))
        fp.TotalLineNumber++
        fp.TotalCharNumber += len(line)
    }
    if err := scanner.Err(); err != nil {
        return Error("Error Scanner")
    }
    fp.Content = content
    return nil
}

func (fp *FileParser) ParseFile() {
    // ParseOneFile
    fp.ParsingCode = true
    stop := false
    for stop != true {
        err := fp.NextLine()
        if err != nil {
            stop = true
        }
    }
}

func (fp *FileParser) NextLine() error {
    fp.ParseLine()
    fp.Index = fp.Index + 1
    if fp.Index >= len(fp.Content) {
        fp.MakeCount()
        fp.State = true
        return Error("End of file")
    }
    return nil
}

func (fp *FileParser) ParseLine() {
    if len(fp.Content) == 0 {
        return
    }
    line := string(fp.Content[fp.Index][:])
    if len(line) == 0 {
        fp.VoidLines++
        //fp.CodeLines--
        return
    }
    if fp.ParsingComment {
        t := fp.Lang.Comment.MultiLineChar
        if fp.Lang.Comment.Inversed {
            t = Reverse(t)
        }
        i := strings.Index(line, t)
        fp.CommentLines++
        if i == -1 {
            fp.CIndex += len(line)
            return
        }
        fp.CIndex += i + len(t)
        fp.MakeCount()
        fp.CIndex += len(line) - (i + len(t))
        return
    }
    fp.CLine = fp.CLine + 1
    i := strings.Index(line, fp.Lang.Comment.OneLineChar)
    if i != -1 {
        if i == 0 {
            fp.CommentLines++
            fp.CLine--
            fp.MakeCount()
            fp.CIndex = len(line)
        } else {
            fp.CIndex += i
            fp.MakeCount()
            fp.CIndex = len(line) - i
        }
        fp.MakeCount()
        return
    }
    i = strings.Index(line, fp.Lang.Comment.MultiLineChar)
    if i != -1 {
        fp.CommentLines++
        if i == 0 {
            fp.CLine--
            fp.MakeCount()
            fp.CIndex += len(line)
        } else {
            fp.CIndex += i
            fp.MakeCount()
            fp.CIndex = len(line) - i
        }
        return
    }
    fp.CIndex += len(line)
    fp.MakeCount()
    fp.MakeCount()
    return
}

func (fp *FileParser) MakeCount() {
    if fp.ParsingCode {
        fp.CodeLenght = fp.CodeLenght + fp.CIndex
        fp.CodeLines = fp.CodeLines + fp.CLine
        fp.CIndex = 0
        fp.CLine = 0
        fp.ParsingCode = false
        fp.ParsingComment = true
    } else {
        fp.CommentLenght = fp.CommentLenght + fp.CIndex
        fp.CommentLines = fp.CommentLines + fp.CLine
        fp.CIndex = 0
        fp.CLine = 0
        fp.ParsingCode = true
        fp.ParsingComment = false
    }
}

func (fp *FileParser) PushArtefacts(f *File) error {
    if fp.State == false {
        return Error("Didnt parse any file")
    }
    f.Stats.TotalLines = fp.TotalLineNumber
    f.Stats.TotalChars = fp.TotalCharNumber
    f.Stats.CodeLines = fp.CodeLines
    f.Stats.CommentLines = fp.CommentLines
    f.Stats.CodeCharsCount = fp.CodeLenght
    f.Stats.CommentCharsCount = fp.CommentLenght
    f.Stats.VoidLines = fp.VoidLines
    return nil
}

func (fp *FileParser) OperateFilePath(filepath string) (File, error) {
    name, ext, lang, err := IdentifyPath(filepath)
    err_load := fp.LoadContent(filepath)
    file := File{
        Stats : Stats{TotalLines : fp.TotalLineNumber, TotalChars : fp.TotalCharNumber},
        Path : filepath,
        Ext : ext,
        Name : name,
        Size : fp.Size,
        ModTime : fp.ModTime,
    }
    if err != nil || err_load != nil {
        return file, Error("File type error")
    }
    file.Lang = lang.Name
    fp.Lang = *lang
    fp.ParseFile()
    fp.PushArtefacts(&file)
    return file, nil
}
