package service

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type InfoFile struct {
	Line string
	File string
}

func ListFiles(dir string, chfiles chan<- []string) []string {
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
	chfiles <- files
	defer close(chfiles)
	return files
}

func (info *InfoFile) ReadFiles(chfiles <-chan []string, chtext chan<- InfoFile) {
	files := <-chfiles
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			//fmt.Println(scanner.Text())
			info.Line = scanner.Text()
			info.File = file
			chtext <- *info
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
	defer close(chtext)

}
