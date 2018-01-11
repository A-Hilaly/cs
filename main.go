package main


import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    path := os.Getenv("PWD")
    if len(os.Args) > 1 {
        path = filepath.Join(path, os.Args[1])
    }
    pd := &Directory{Path : path}
    _, _ = pd.WalkAndWork(true, LoadGitIgnore(path).List)

    t := &Tree{
        Head : pd,
    }
    t.Repr()
    //fmt.Println("[Files] :", nf)
    //fmt.Println("[Directories] :", nd)
    //fmt.Println("[size]", t.Head.Size)
    fmt.Println("Stats :")
    t.Head.CollectStats()
    t.Head.Stats.Show()
}
