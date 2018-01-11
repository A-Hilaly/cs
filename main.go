package main


import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    path := os.Getenv("PWD")
    if len(os.Args) > 1 && os.Args[0] != "." {
        path = filepath.Join(path, os.Args[1])
    }
    pd := &Directory{Path : path}
    _, _ = pd.WalkAndWork(false, LoadGitIgnore(path).List)

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
