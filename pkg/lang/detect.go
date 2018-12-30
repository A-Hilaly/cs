package lang

import "path/filepath"

var (
	langArray = []*Spec{
		C, CPP, Golang,
		JS, TS, CSharp,
		Bash, Yaml,
	}
)

type Detector interface {
	Detect(string) (*Spec, error)
}

func NewDetector() *detector {
	return &detector{}
}

type detector struct{}

func (d *detector) Detect(filename string) (*Spec, error) {
	ext := filepath.Ext(filename)
	for _, lang := range langArray {
		for _, ex := range lang.Extensions {
			if ext == ex {
				return lang, nil
			}
		}
	}
	return UnknownLang, nil
}
