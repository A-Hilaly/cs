package main

import "time"

type File struct {
    Stats
    ParentDirectory   *Directory
    Path              string
    Ext               string
    Name              string
    ModTime           time.Time
    Lang              string
    Size              int64
}

func (file *File) GenerationsAbove() int {
    if file.ParentDirectory == nil {
        return 0
    }
    return file.ParentDirectory.GenerationsAbove() + 1
}
