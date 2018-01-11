package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "time"
)


type File struct {
    ParentDirectory   *Directory
    Path              string
    Ext               string
    Name              string
    ModTime           time.Time
    Lang              string
    Size              int64
    TotalLines        int
    TotalChars        int
    CodeLines         int
    CommentLines      int
    CodeCharsCount    int
    CommentCharsCount int
    VoidLines         int
}

func (file *File) IncrSize(size int) {
    file.Size += int64(size)
    if file.ParentDirectory != nil {
        file.ParentDirectory.IncrSize(size)
    }
}

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
    if err != nil {
        //log.Fatal(err)
        return Error("Open error")
    }
    err = fp.LoadStats(file)
    if err != nil {
        return Error("Stats Error")
    }
    content := make([][]byte, 0)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        content = append(content, []byte(line))
        fp.TotalLineNumber++
        fp.TotalCharNumber += len(line)
        //fmt.Println("Reading : ", content)
    }
    if err := scanner.Err(); err != nil {
        //fmt.Println("Error Scanner")
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
    //fmt.Println("INFO : ", len(fp.Content), fp.Index)
    fp.ParseLine()
    fp.Index = fp.Index + 1
    if fp.Index == len(fp.Content) {
        fp.MakeCount()
        fp.State = true
        return Error("End of file")
    }
    return nil
}

func (fp *FileParser) ParseLine() {
    // ParsingCode should be true for first line
    line := string(fp.Content[fp.Index][:])
    if len(line) == 0 {
        fp.VoidLines++
        //fp.CodeLines--
        return
    }
    if fp.ParsingComment {
        i := strings.Index(line, fp.Lang.Comment.MultiLineChar)
        fp.CommentLines++
        if i == -1 {
            fp.CIndex += len(line)
            return
        }
        fp.CIndex += i + len(fp.Lang.Comment.MultiLineChar)
        //fmt.Println("---", fp.CIndex)
        fp.MakeCount()
        fp.CIndex += len(line) - (i + len(fp.Lang.Comment.MultiLineChar))
        //fp.CIndex += len(fp.Lang.Comment.MultiLineChar) + 1 - i
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
    f.TotalLines = fp.TotalLineNumber
    f.TotalChars = fp.TotalCharNumber
    f.CodeLines = fp.CodeLines
    f.CommentLines = fp.CommentLines
    f.CodeCharsCount = fp.CodeLenght
    f.CommentCharsCount = fp.CommentLenght
    f.VoidLines = fp.VoidLines
    return nil
}

func (fp *FileParser) OperateFilePath(filepath string) (File, error) {
    name, ext, lang, err := IdentifyPath(filepath)
    err_load := fp.LoadContent(filepath)
    file := File{
        Path : filepath,
        Ext : ext,
        Name : name,
        Size : fp.Size,
        ModTime : fp.ModTime,
        TotalLines : fp.TotalLineNumber,
        TotalChars : fp.TotalCharNumber,
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

func main() {
    //fmt.Println(Python.Comment.Inversed)
    //a := FileParser{}
    //file, err := a.OperateFilePath("test.py")
    //fmt.Println(file, err)
    pd := &Directory{Path : os.Getenv("PWD") + "/dodo"}
    pd.WalkAndWork()
    //fmt.Println(pd.Path, pd.Name)
    fmt.Println(pd.SubDirectories[0].SubFiles[0].GenerationsAbove())

}
