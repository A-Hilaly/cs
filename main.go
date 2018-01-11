package main


import (
    "os"
    "fmt"
    "time"
    "path/filepath"
    "github.com/a-hilaly/gocs/gocs"
)

func main() {
    start := time.Now()
    path := os.Getenv("PWD")
    if len(os.Args) > 1 {
        if os.Args[0] != "." {
            path = filepath.Join(path, os.Args[1])
        }
    } else {
        fmt.Println("-h --help")
        return
    }
    pd := &gocs.Directory{Path : path}
    nd, nf := pd.WalkAndWork(true, gocs.LoadGitIgnore(path).List)
    tm := timeTrack(start, "Walk And Work")
    t := &gocs.Tree{
        Head : pd,
    }
    t.Repr()
    tm = timeTrack(tm, "Repr Tree")
    fmt.Println("[Files] :", nf)
    fmt.Println("[Directories] :", nd)
    fmt.Println("[size]", t.Head.Size)
    fmt.Println("Stats :")
    t.Head.CollectStats()
    tm = timeTrack(tm, "Collect Stats")
    t.Head.Stats.Show()
    tm = timeTrack(start, "Total")
}

func timeTrack(start time.Time, name string) time.Time {
    elapsed := time.Since(start)
    fmt.Println("STEP ", name, "<>", elapsed)
    return time.Now()
}
