package main

import (
    "path/filepath"
    "os"
)

type WalkMan struct {
    StartPath       string
    IgnoreRegexList []string
    Dirs            []string
    Files           []string
    TreeStruct      Directory
}

type Walker interface {
    Walk(path string) (err error)
}

func (w *WalkMan) ValidateStartingPath() error {
    // Validate Start Path
    if w.StartPath == "" {
        return Error("Empty starting path")
    }
    if _, err := os.Stat(w.StartPath); err != nil {
        return Error("Starting path doesnt exists")
    }
    return nil
}

func (w *WalkMan) Walk() (err error) {
    // return list of files/dirs inside directory StartPath
    err = filepath.Walk(w.StartPath, func (path string, f os.FileInfo, err error) error {
        if f.IsDir() {
            w.Dirs = append(w.Dirs, path)
        } else {
            w.Files = append(w.Files, path)
        }
        return nil
    })
    return
}

func (w *WalkMan) GetLanguageFiles(lang string) (list []string, err error) {
    ext, err := IdentifyLanguage(lang)
    if err != nil {
        return nil, err
    }
    for _, elem := range w.Files {
        if filepath.Ext(elem) == ext {
            list = append(list, elem)
        }
    }
    return list, nil
}


func (w *WalkMan) AnalyseFileCases() (slist []string, unknown []string) {
    for _, elem := range w.Files {
        ext := filepath.Ext(elem)
        if _, err := IdentifyExtention(ext); err == nil {
            slist = append(list, elem)
        } else {
            unknown = append(list, elem)
        }
    }
    return list
}
