package service

import (
	"io/fs"
	"log"
	"path/filepath"
)

func ListFiles(dir string) []string {
    var files []string

    err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
       if !d.IsDir() /* && filepath.Ext(path) == ".md" */ {
          files = append(files, path)
       }
       return nil
    })
    if err != nil {
       log.Fatal(err)
    }

    return files
}

