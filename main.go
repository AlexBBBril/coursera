package main

import (
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

type Dir struct {
	Name string
}

type Structure []Dir
func (a Structure) Len() int           { return len(a) }
func (a Structure) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Structure) Less(i, j int) bool { return a[i].Name < a[j].Name }

func (p Dir) String() string {
	return ``
}

func dirTree(in io.Writer, path string, printFiles bool) error {
	file, err := os.Open(path)
	if nil != err {
		return err
	}

	fileInfo, err := file.Readdir(0)
	if nil != err {
		return err
	}

	var structure = make(Structure, 0)

	for _, info := range fileInfo {
		if info.IsDir() {
			ff, err := os.Open(filepath.Join(path, info.Name()))
			if nil != err {
				return err
			}

			fi, err := ff.Readdir(0)
			if nil != err {
				return err
			}
			_ = fi

			structure = append(structure, Dir{Name: info.Name()})
		}

	}

	sort.Sort(structure)
	strings.Contains(path, ".")

	filepath.Abs(path)

	return nil
}

func isDir(path string) error {
	file, err := os.Open(path)
	if nil != err {
		return err
	}

	fileInfo, err := file.Readdir(0)
	if nil != err {
		return err
	}

	for _, info := range fileInfo {
		if info.IsDir() {
			ff, err := os.Open(filepath.Join(path, info.Name()))
			if nil != err {
				return err
			}

			fi, err := ff.Readdir(0)
			if nil != err {
				return err
			}
			_ = fi

			structure = append(structure, Dir{Name: info.Name()})
		}

	}
}
