package main

import (
    "os"
    //"fmt"
    "regexp"
    "io/ioutil"
    "path/filepath"
)

type Directory struct {
    Stats
    ParentDirectory *Directory
    SubDirectories  []*Directory
    SubFiles        []*File
    Path            string
    Name            string
    Size            int64
}

func (dir *Directory) Exists() bool {
    if _, err := os.Stat(dir.Path); err != nil {
        return false
    }
    return true
}

func (dir *Directory) CollectStats() {
    for _, file := range dir.SubFiles {
        dir.Stats.AppendFromStats(file.Stats)
    }
    for _, d := range dir.SubDirectories {
        d.CollectStats()
        dir.Stats.AppendFromStats(d.Stats)
    }
}

func (dir *Directory) GenerationsAbove() int {
    if dir.ParentDirectory == nil {
        return 0
    }
    return dir.ParentDirectory.GenerationsAbove() + 1
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

func (dir *Directory) SelfResolve() error {
    dir.Name = filepath.Base(dir.Path)
    file, err := OpenFile(dir.Path)
    defer file.Close()
    if err != nil {
        return err
    }
    stats, err := file.Stat()
    if err != nil {
        return err
    }
    dir.Size = stats.Size()
    return nil
}

func Ignore(target string, l []string) bool {
    for _, elem := range l {
        if re := regexp.MustCompile(elem); re.MatchString(target) {
            return true
        }
    }
    return false
}

func (dir *Directory) WalkAndWork(regignore []string) (int, int) {
    _ = dir.SelfResolve()
    d, _ := ioutil.ReadDir(dir.Path)
    nfiles, ndirs := 0, 0 // number of files // number of dir
    for _, elem := range d {
        if Ignore(elem.Name(), regignore) {
            //fmt.Println("True", elem.Name())
            continue
        }
        if elem.IsDir() {
            dr := &Directory{
                Path : filepath.Join(dir.Path, elem.Name()),
                Size : elem.Size(),
                Name : elem.Name(),
            }
            dir.AppendDirectory(dr)
            //fmt.Println(dr)
            nfiles += 1
            nf, nd := dr.WalkAndWork(regignore)
            nfiles, ndirs = nfiles + nf, ndirs + nd
        } else {
            fp := FileParser{}
            file, _ := fp.OperateFilePath(filepath.Join(dir.Path, elem.Name()))
            //fmt.Println(file)
            ndirs += 1
            dir.AppendFile(&file)
        }
    }
    return nfiles, ndirs
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
