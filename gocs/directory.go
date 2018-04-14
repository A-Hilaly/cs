package gocs

import (
    "os"
    "regexp"
    "io/ioutil"
    "path/filepath"
)
type GitIgnore struct {
    List  []string
}

func LoadGitIgnore(path string) *GitIgnore {
    _, err := OpenFile(path + "/.gitignore")
    if err != nil {
        return &GitIgnore{}
    } else {
        fileparser := FileParser{}
        err = fileparser.LoadContent(path + "/.gitignore")
        if err != nil {
            return &GitIgnore{}
        }
        l := []string{}
        for _, elem := range fileparser.Content {
            if len(elem) > 0 && string(elem[0]) != "#" {
                l = append(l, "^(." + string(elem) + ")")
            }
        }
        return &GitIgnore{
            List : l,
        }
    }
}

type Directory struct {
    Stats
    ParentDirectory *Directory
    SubDirectories  []*Directory
    SubFiles        []*File
    Path            string
    Name            string
    Size            int64
}

var fp = FileParser{}

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
        if re := regexp.MustCompile(elem); elem == target || re.MatchString(target) {
            //fmt.Println(target, elem, "hey")
            return true
        }
    }
    return false
}

func (dir *Directory) WalkAndWork(ih bool, gil []string) (int, int) {
    _ = dir.SelfResolve()
    d, _ := ioutil.ReadDir(dir.Path)
    regignore := []string{}
    regignore = append(regignore, "^(.git)$")
    if ih {
        regignore = append(regignore, "^[.]")
    }
    regignore = append(regignore, gil...)
    nfiles, ndirs := 0, 0 // number of files // number of dir
    for _, elem := range d {
        if Ignore(elem.Name(), regignore) {
            //fmt.Println(elem.Name())
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
            nf, nd := dr.WalkAndWork(ih, gil)
            nfiles, ndirs = nfiles + nf, ndirs + nd
        } else {
            file, _ := fp.OperateFilePath(filepath.Join(dir.Path, elem.Name()))
            fp = FileParser{}
            //fmt.Println(file)
            ndirs += 1
            dir.AppendFile(&file)
        }
    }
    return nfiles, ndirs
}


func WorkConcurent(dir *Directory, ih bool, gil []string) (chan Stats, chan int, chan int) {
    _ = dir.SelfResolve()

    regignore := []string{}
    regignore = append(regignore, "^(.git)$")
    if ih {
        regignore = append(regignore, "^[.]")
    }
    regignore = append(regignore, gil...)
    nfiles, ndirs := 0, 0 // number of files // number of dir
    schan, nc, np := make(chan Stats, 1), make(chan int, 1), make(chan int, 1)
    go GetStats(dir.Path, reignore, schan, nc, np)
    return schan, nc, np
}

func GetStats(dir string, regignore []string, schan chan Stats, nc chan int, np chan int) {
    totalstats := Stats{}
    allstats := []Stats

    vs := make(chan File)
    done := make(chan bool, 1)

    nf, nd := 0, 0

    d, _ := ioutil.ReadDir(dir)

    for _, elem := range d {
        if Ignore(elem.Name(), regignore) {
            //fmt.Println(elem.Name())
            continue
        }
        if elem.IsDir() {
            //fmt.Println(dr)
            nf += 1
            nf, nd := WalkAndWork(&dr, ih, gil)
            nfiles, ndirs = nfiles + nf, ndirs + nd
        } else {
            file, _ := fp.OperateFilePath(filepath.Join(dir.Path, elem.Name()))
            fp = FileParser{}
            //fmt.Println(file)
            ndirs += 1
            dir.AppendFile(&file)
        }
    }

    for {
        select{
        case <- done:
            schan <- totalstats
            nc <- nf
            np <- nd
            return
        }
    }
}
