package main

import "fmt"
import "io/ioutil"
import "path/filepath"

type Directory struct {
    ParentDirectory *Directory
    SubDirectories  []*Directory
    SubFiles        []*File
    Path            string
    Name            string
    TotalSize       int64
}

type Tree struct {
    Head *Directory
}

func (tree *Tree) Repr() {
    tree.Head.Repr()
}

func (dir *Directory) Repr() {
    return
}

func (dir *Directory) GenerationsAbove() int {
    if dir.ParentDirectory == nil {
        return 0
    }
    return dir.ParentDirectory.GenerationsAbove() + 1
}

func (file *File) GenerationsAbove() int {
    if file.ParentDirectory == nil {
        return 0
    }
    return file.ParentDirectory.GenerationsAbove() + 1
}

func (dir *Directory) IncrSize(size int) {
    dir.TotalSize += int64(size)
    if dir.ParentDirectory != nil {
        dir.ParentDirectory.IncrSize(size)
    }
}

func (dir *Directory) AppendDirectory(d *Directory) {
    dir.SubDirectories = append(dir.SubDirectories, d)
    d.ParentDirectory = dir
}

func (dir *Directory) AppendFile(f *File) {
    dir.SubFiles = append(dir.SubFiles, f)
    f.ParentDirectory = dir
}

func (dir *Directory) SimpleWalk() {
    d, _ := ioutil.ReadDir(dir.Path)
    for _, elem := range d {
        if elem.IsDir() {
            dr := &Directory{Path : filepath.Join(dir.Path, elem.Name())}
            dir.AppendDirectory(dr)
            dr.SimpleWalk()
        } else {
            f := &File{Path : filepath.Join(dir.Path, elem.Name())}
            dir.AppendFile(f)
        }
    }
}

func (dir *Directory) WalkAndWork() {
    d, _ := ioutil.ReadDir(dir.Path)
    dir.Name = filepath.Base(dir.Path)
    for _, elem := range d {
        if elem.IsDir() {
            dr := &Directory{
                Path : filepath.Join(dir.Path, elem.Name()),
                TotalSize : elem.Size(),
                Name      : elem.Name(),
            }
            dir.AppendDirectory(dr)
            dr.WalkAndWork()
        } else {
            fp := FileParser{}
            file, _ := fp.OperateFilePath(filepath.Join(dir.Path, elem.Name()))
            fmt.Println(file)
            dir.AppendFile(&file)
        }
    }
}

/*
func main() {
    pd := &Directory{}
    d1 := &Directory{}
    pd.AppendDirectory(d1)
    fmt.Println(pd, d1)
    d1.IncrSize(1)
    fmt.Println(pd, d1)
}
*/
