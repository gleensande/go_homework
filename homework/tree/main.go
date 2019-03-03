package main

import (
	"fmt"
//	"io"
	"os"
	"path/filepath"
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

func dirTree(out *os.File, path string, printFiles bool) error {
	err := filepath.Walk(path, func(pathToFile string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if pathToFile == path {
			return nil
		}

		for i := 0; i < strings.Count(pathToFile, string(os.PathSeparator)); i++ {
			fmt.Print("    ")
		}
		if info.IsDir() {
			fmt.Print("└───")
		} else {
			fmt.Print("├───")
		}
		fmt.Print(info.Name(), "\n")

		return nil
	})
	return err
}
