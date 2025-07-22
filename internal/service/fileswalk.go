package service

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

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

func ReadFiles(chfiles <-chan []string, chtext chan<- string) {
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
			chtext <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
	defer close(chtext)

}
