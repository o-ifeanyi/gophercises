package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path, err := filepath.Abs("sample")
	if err != nil {
		log.Println(err)
	}
	err = filepath.WalkDir(path, rename)
	if err != nil {
		log.Println("Something went wrong:", err)
	}
}

func rename(path string, d fs.DirEntry, err error) error {
	name := d.Name()
	if !d.IsDir() && strings.Contains(name, "_") {
		paths := strings.Split(path, "/")
		res := strings.Split(name, ".")
		filename := res[0]
		res = strings.Split(filename, "_")

		res[1] = fmt.Sprintf("(%s of 100).txt", strings.Trim(res[1], "0"))
		paths[len(paths)-1] = strings.Join(res, " ")
		fmt.Println("renaming", path, "to", strings.Join(paths, "/"))
		err = os.Rename(path, strings.Join(paths, "/"))
	}
	return err
}
