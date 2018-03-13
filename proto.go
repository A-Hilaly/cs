package main

import (
    "fmt"
    "time"
    "github.com/a-hilaly/cs/gocs"
)

func FileParserWorker(p string, result chan gocs.File) {
    fp := gocs.FileParser{}
    file, _ := fp.OperateFilePath(p)
    result <- file
}

func main() {
    t := make(chan gocs.File, 1)
    go FileParserWorker("main.go", t)
    fmt.Println("Hello")
    time.Sleep(time.Second * 1)
    fmt.Println(<-t)
}
