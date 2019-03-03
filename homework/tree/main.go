package main

import (
//	"fmt"
	"io"
	"io/ioutil"
	"os"
//	"path/filepath"
	"strings"
	"strconv"
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

func dirTree(out io.Writer, path string, printFiles bool) error {
	levels := make(map[int]bool)

	err := watchDir(out, path, printFiles, levels)

	return err
}

func watchDir(out io.Writer, path string, printFiles bool, levels map[int]bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	if !printFiles {
		files = dirsOnly(files)
	}

	numOfFiles := len(files)

	for i := 0; i < numOfFiles; i++ {
		pathToFile := path + string(os.PathSeparator) + files[i].Name()
		level := strings.Count(pathToFile, string(os.PathSeparator))

		for j := 1; j < level; j++ {
			if levels[j] {
				out.Write([]byte("│"))
			}
			out.Write([]byte("\t"))
		}

		if i != numOfFiles - 1 {
			out.Write([]byte("├───"))
			levels[level] = true
		} else {
			out.Write([]byte("└───"))
			delete(levels, level)
		}
		out.Write([]byte(files[i].Name()))

		if files[i].IsDir() {
			out.Write([]byte("\n"))
			watchDir(out, pathToFile, printFiles, levels)
		} else {
			out.Write(sizeOfFile(files[i].Size()))
			out.Write([]byte("\n"))
		}
	}

	return err
}

func sizeOfFile(size int64) []byte {
	var result string
	if size == 0 {
		result = " (empty)"
	} else {
		result = " (" + strconv.Itoa(int(size)) + "b)"
	}

	return []byte(result)
}

func dirsOnly(files []os.FileInfo) []os.FileInfo {
	var directories []os.FileInfo
	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file)
		}
	}
	return directories
}
