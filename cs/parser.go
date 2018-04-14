package cs

import (
    b8 "bytes"
    "errors"
    "io/ioutil"
    "path/filepath"
)
type FileParser interface {
    Parse(path string) (Stats, error)
}

type ConcurrentParser interface {
    ParseConcurrent(ts chan Stats, max_threads int, paths ...string) (Stats, error)
}

type CsParser struct {
    Type string
}

const (
    endl byte = '\n'
)

func parseFile(path string) (*Stats, error) {
    //
    ext  := filepath.Base(path)
    base := filepath.Ext(path)
    bytes, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, errors.New("File not found")
    }
    lang, err := getLanguageSpecs(ext)
    if err != nil {
        return &Stats{
                Language   : unknownLang(ext),
                Size       : len(bytes),
                TotalChars : len(bytes),
                TotalLines : b8.Count(bytes, []byte{endl})
        }, nil
    }
    stats := Stats{
        Language : lang,
        Size     : len(bytes),
    }
    // comment range OL/ML
    crol, crml := lang.Comment.OL, lang.Comment.ML

    //Total lines
    //Total Char
    //Code  Lines
    //Comments Lines
    //Code chars count
    //Comment chars count
    //Void Lines
    tl, tc, cdl, cml, cmcc, cdcc, vl := 0, 0, 0, 0, 0, 0, 0
    lang_cr    := len(crol)
    lang_mcr   := len(crml)
    is_code    := true
    is_comment := false
    line_index := 0

    for i, b := range bytes {
        if b != endl {
            line_index++
            continue
        }
        if line_index == 0 {
            tl++
            vl++
            if is_code {
                cdl++
            } else {
                cml++
            }
            continue
        }
        tl++
        if bytes[i:line_index] {

        }
        if b == {

        }

        line_index = 0
    }
}

func (csp *CsParser) Parse(path string) (Stats, error) {
    return parseFile(path)
}


func (csp *CsParser) ParseConcurrent(ts chan Stats, max_threads int, paths ...string) {

}

func (csp *CsParser) SmartConcurrent(ts chan Stats, max_threads int, path string) {
    fstats  := Stats{}
    defer ts <- fstats
    files, dirs := WalkOver(path)
    nf, nd := len(files), len(dirs)
    fs := make(chan Stats, nf / 8)

    mthreads := len(files) / max_threads
    ethreads := len(files) % max_threads
    all_chans := make([]chan Stats, 100)
    for i := 0; i <= mthreads + ethreads; i++ {
        go csp.ParseConcurrent(all_chans[i], max_threads, files[i, i+max_threads]...)
    }
}
