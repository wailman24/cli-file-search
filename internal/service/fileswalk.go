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
	NumL int
}

func ListFiles(dir string, chfiles chan<- []string, ext string, ignore string) []string {
	var files []string
	if dir == "" {
		return nil
	}

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		/* var fignore, dignore string

		if ignore != "" && ignore[0] == '.' {
			fignore = ignore
		} else {
			dignore = ignore
		} */

		if !d.IsDir() && (filepath.Ext(path) == ext || ext == "") && (filepath.Ext(path) != ignore || ignore == "") {
			files = append(files, path)
		} else if d.IsDir() && (d.Name() == ignore) {
			return filepath.SkipDir
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

		scanner := bufio.NewScanner(f)
		i := 0
		for scanner.Scan() {
			i++
			//fmt.Println(scanner.Text())
			if scanner.Text() != "" {
				info.NumL = i
				info.Line = scanner.Text()
				info.File = file
				chtext <- *info
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		f.Close()
	}

	defer close(chtext)

}
