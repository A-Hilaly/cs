# GoCS
Code statistics command line

[![CircleCI](https://circleci.com/gh/A-Hilaly/cs.svg?style=svg)](https://circleci.com/gh/A-Hilaly/cs)

###### Feature 

languages: C, C++, C#, Go, Javascript, typescript, Bash, Yaml

---

#### Install

```shell
go get -u github.com/a-hilaly/cs/...
```

#### Usage

```bash
code statistics

Usage:
  cs [flags]

Flags:
  -h, --help            help for cs
  -g, --is-git          is true will lead program to ignore all files under .git directory (default true)
  -G, --use-gitignore   i .git directory (default true)
```

examples:

```bash
$ cs test/testdata
+----------+-------+-------+------+---------+-------+
| LANGUAGE | FILES | TOTAL | CODE | COMMENT | BLANK |
+----------+-------+-------+------+---------+-------+
| Bash     |     1 |     4 |    1 |       3 |     0 |
| C        |     1 |    49 |   32 |      18 |     8 |
| Cpp      |     1 |    42 |   29 |      13 |    11 |
| Go       |     3 |    27 |   16 |      12 |    10 |
| Yaml     |     1 |     6 |    4 |       3 |     0 |
+----------+-------+-------+------+---------+-------+
|    TOTAL |     7 |   128 |   82 |      49 |    29 |
+----------+-------+-------+------+---------+-------+
```

#### Todo

- Ignore .git directory
- Use .gitignore elements
- Add unit tests
- Add benchmark tests

support langages:
- ocaml
- Ruby
- css
- html/xml
- rust
- python
- Makefile