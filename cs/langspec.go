package cs

import (
    "sync"
    "errors"
)

type Comment struct {
    OL       string
    ML       string
    Inversed bool
}

type Language struct {
    Name      string
    Extension string
    Comment   Comment
}

const (
    Python = Language{
        Name : "Python",
        Extension : ".py",
        Comment : Comment{
            OL : "#",
            ML : "\"\"\"",
            Inversed : false,
        },
    }
    Golang = Language{
        Name : "Golang",
        Extension : ".go",
        Comment : Comment{
            OL : "//",
            ML : "/*",
            Inversed : true,
        },
    }
    Unknown = Language{
        Name : "UnknownLanguage",
        Extention : "",
    }
)

var support = []Language{Python, Golang}

func getLanguageSpecs(ext string) (*Language, error) {
    for _, elem := range support {
        if elem.Extention == ext {
            return &elem, nil
        }
    }
    return nil, errors.New("Unknown language")
}

var cached_unknown = []*Language{}

var cache_mutex = sync.Mutex{}

func unknownLang(ext string) *Language {
    cached_mutex.Lock()
    defer cached_mutex.Unlock()
    for _, elem := range cached_unknown {
        if elem.Extension == ext {
            return &elem
        }
    }
    lang := &Language{
        Name : "UnknownLanguage",
        Extention : ext,
    }
    append(cached_unknown, lang)
    return lang
}
