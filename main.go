package main


import (
    "os"
    "fmt"
    "time"
    "path/filepath"
)

func timeTrack(start time.Time, name string) time.Time {
    elapsed := time.Since(start)
    fmt.Println("::::  STEP :", name, elapsed)
    return time.Now()
}

func main() {
    start := time.Now()
    path := os.Getenv("PWD")
    if len(os.Args) > 1 && os.Args[0] != "." {
        path = filepath.Join(path, os.Args[1])
    }
    tm := timeTrack(start, "Args Treating")
    pd := &Directory{Path : path}
    nd, nf := pd.WalkAndWork(false, LoadGitIgnore(path).List)
    tm = timeTrack(tm, "Walk And Work")
    t := &Tree{
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
    tm = timeTrack(tm, "Total")
}
