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
