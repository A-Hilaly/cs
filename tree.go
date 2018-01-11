package main

import (
    "fmt"
    "strings"
)


type Tree struct {
    Head *Directory
}

func (tree *Tree) Repr() {
    tree.Head.ReprDir()
    tree.Head.Repr()
}

func (dir *Directory) Repr() {
    if len(dir.SubDirectories) != 0 {
        for _, elem := range dir.SubDirectories {
            elem.ReprDir()
            elem.Repr()
        }
    }
    dir.ReprFileGeneration()
}

func (dir *Directory) ReprFileGeneration() {
    c := dir.GenerationsAbove()
    //fmt.Println(dir.Name, "have ", c, "generations")
    indent := ""
    if c > 0 {
        indent = strings.Repeat("│   ", c)
    }
    for i, elem := range dir.SubFiles {
        if i == len(dir.SubFiles) - 1{
            fmt.Println(indent + "└── " + elem.Name)
        } else {
			fmt.Println(indent + "├── " + elem.Name)
		}
    }
}

func (dir *Directory) ReprDir() {
    c := dir.GenerationsAbove()
    if c == 0 {
        fmt.Println(dir.Name)
    } else {
        indent := strings.Repeat("│   ", c - 1) + "├── "
        fmt.Println(indent + dir.Name)
    }
}

/*
func (dir *Directory) ReprDirsGeneration() {
    c := dir.GenerationsAbove()
    indent := ""
    if c > 1 {
        indent = strings.Repeat("│   ", c)
    }
    for i, elem := range dir.SubDirectories {
        if i == len(dir.SubDirectories) - 1 {
            fmt.Println(indent + "└── " + elem.Name)
        } else {
			fmt.Println(indent + "├── " + elem.Name)
		}
    }
}
*/
