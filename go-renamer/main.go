package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type rename struct {
	dir      string
	folder   string
	filename string
}

func main() {
	path, err := filepath.Abs("sample")
	if err != nil {
		log.Println(err)
	}
	toRename := make(map[string][]rename)
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		name := d.Name()
		if _, ok := toRename[name]; !ok && d.IsDir() {
			toRename[name] = []rename{}
			return nil
		}
		if !d.IsDir() && strings.Contains(name, "_") {
			dir := filepath.Dir(path)
			for k, v := range toRename {
				if strings.HasSuffix(dir, k) {
					toRename[k] = append(v, rename{
						dir: dir, folder: k, filename: name,
					})
				}
			}
		}
		return err
	})
	if err != nil {
		log.Println("Something went wrong:", err)
	}

	for _, v := range toRename {
		renameFile(v)
	}
}

func renameFile(re []rename) (err error) {
	for i, r := range re {
		res := strings.Split(r.filename, ".")
		filename := res[0]
		res = strings.Split(filename, "_")
		res[1] = fmt.Sprintf("(%d of %d)%s", i+1, len(re), filepath.Ext(r.filename))
		newFilename := strings.Join(res, " ")

		err = os.Rename(filepath.Join(r.dir, r.filename), filepath.Join(r.dir, newFilename))
		if err != nil {
			return err
		}
	}
	return err
}
