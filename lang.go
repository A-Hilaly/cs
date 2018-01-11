package main

import "path/filepath"

type Comment struct {
    OneLineChar   string
    MultiLineChar string
    Inversed      bool
}

type Language struct {
    Name          string
    Extension     string
    Comment       *Comment
}

var (
    PythonComment = &Comment{
        OneLineChar : "#",
        MultiLineChar : "\"\"\"",
        Inversed : false,
    }
    CommentGolang = &Comment{
        OneLineChar : "//",
        MultiLineChar : "/*",
        Inversed : true,
    }
)

var (
    Python = &Language{
        Name : "Python",
        Extension : ".py",
        Comment : PythonComment,
    }
)

var (
    Golang = &Language{
        Name : "Golang",
        Extension : ".go",
        Comment : CommentGolang,
    }
)

var LANG_SUPPORT = [2]*Language{Python, Golang}

func IdentifyLanguage(lang string) (*Language, error) {
    for _, elem := range LANG_SUPPORT {
        if elem.Name == lang {
            return elem, nil
        }
    }
    return nil, Error("Language is not supported")
}

func IdentifyExtention(ext string) (*Language, error) {
    for _, elem := range LANG_SUPPORT {
        if elem.Extension == ext {
            return elem, nil
        }
    }
    return nil, Error("Language is not supported")
}

func IdentifyPath(path string) (name string, ext string, lang *Language, err error) {
    ext = filepath.Ext(path)
    name = filepath.Base(path)
    lang, err = IdentifyExtention(ext)
    return name, ext, lang, err
}
