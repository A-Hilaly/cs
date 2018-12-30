package walker

import (
	"fmt"
	"io/ioutil"

	"github.com/a-hilaly/cs/pkg/lang"
	"github.com/a-hilaly/cs/pkg/parser"
	"github.com/a-hilaly/cs/pkg/stats"
)

const (
	GIT       = ".git"
	GITIGNORE = ".gitignore"
)

type Walker interface {
	Walk(path string) (*stats.Total, error)
}

type walker struct {
	config *Config
	ld     lang.Detector
	parser parser.Parser
}

func New(c *Config) *walker {
	return &walker{
		config: c,
		ld:     lang.NewDetector(),
		parser: parser.ParseFileBytes,
	}
}

func (w *walker) Walk(path string) (*stats.Total, error) {
	if !w.config.CC {
		return w.simpleWalk(path)
	} else {
		return w.concurrentWalk(path)
	}
}

func (w *walker) simpleWalk(path string) (*stats.Total, error) {
	files, err := resolveFilePaths(path)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve path: %v", path)
	}

	sts := &stats.Total{
		StatsPerLang: make(map[string]*stats.Total, 0),
	}
	for _, file := range files {
		lang, _ := w.ld.Detect(file)

		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("reading file: %v", bytes)
		}

		s, err := w.parser(bytes, lang)
		if err != nil {
			return nil, fmt.Errorf("parsing file: %v", err)
		}

		if _, ok := sts.StatsPerLang[lang.Name]; !ok {
			sts.StatsPerLang[lang.Name] = &stats.Total{
				StatsPerLang: make(map[string]*stats.Total, 0),
			}
		}
		sts.Append(s)
	}
	return sts, nil
}

func (w *walker) concurrentWalk(path string) (*stats.Total, error) {
	return nil, fmt.Errorf("not implemented")
}
