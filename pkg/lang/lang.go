package lang

type CommentStyle struct {
	One        []byte
	BlockStart []byte
	BlockEnd   []byte
}

var (
	CComment = &CommentStyle{
		One:        []byte("//"),
		BlockStart: []byte("/*"),
		BlockEnd:   []byte("*/"),
	}

	BashComment = &CommentStyle{
		One: []byte("#"),
	}

	PythonComment = &CommentStyle{
		One:        []byte("#"),
		BlockStart: []byte("\"\"\""),
		BlockEnd:   []byte("\"\"\""),
	}
)

type Spec struct {
	Name       string
	Extensions []string
	Comment    *CommentStyle
}

var (
	// C like languages
	C = &Spec{
		Name:       "C",
		Extensions: []string{".c", ".h"},
		Comment:    CComment,
	}

	Golang = &Spec{
		Name:       "Go",
		Extensions: []string{".go"},
		Comment:    CComment,
	}

	CPP = &Spec{
		Name:       "Cpp",
		Extensions: []string{".cpp", ".c++", ".hpp"},
		Comment:    CComment,
	}

	CSharp = &Spec{
		Name:       "CSharp",
		Extensions: []string{".cs", ".csharp"},
		Comment:    CComment,
	}

	JS = &Spec{
		Name:       "Javascript",
		Extensions: []string{".js"},
		Comment:    CComment,
	}

	TS = &Spec{
		Name:       "Typescript",
		Extensions: []string{".ts"},
		Comment:    CComment,
	}

	// XML like languages
	XML  = &Spec{}
	HTML = XML

	// Bash like languages
	Bash = &Spec{
		Name:       "Bash",
		Extensions: []string{".sh"},
		Comment:    BashComment,
	}

	Yaml = &Spec{
		Name:       "Yaml",
		Extensions: []string{".yml", ".yaml"},
		Comment:    BashComment,
	}

	Makefile   = Bash
	Dockerfile = Bash

	// Python
	Python = &Spec{}

	UnknownLang = &Spec{
		Name: "Unknown",
	}
)
