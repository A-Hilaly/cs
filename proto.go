package main

import (
    "fmt"
)

type Stats struct {
    Path string
}

func ParseFile(path string) (Stats, error) {
    return Stats{path}, nil
}

func WorkerParser(path string, sts chan Stats) {
    stats, _ := ParseFile(path)
    sts <- stats
}

func MultiWorkerParser(paths ...string) {
    lp := len(paths)
    fmt.Println(lp)
    nice_channel := make(chan Stats, lp)
    for i := 0; i < lp; i++ {
        go WorkerParser(paths[i], nice_channel)
    }
    count := 0
    for {
        select {
        case m := <- nice_channel:
            fmt.Println(m)
            count++
            if count == lp {
                return
            }
        }

    }
}

func main() {
    MultiWorkerParser([]string{"main.go", "proto.go", ".gitignore"}...)
}
