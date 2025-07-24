package service

import (
	"bufio"
	"fmt"
	"io"
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
	var isbinary bool
	if dir == "" {
		return nil
	}

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		isbinary, _ = IsBinaryFile(path)

		if !isbinary && !d.IsDir() && (filepath.Ext(path) == ext || ext == "") && (filepath.Ext(path) != ignore || ignore == "") {
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

func IsBinaryFile(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return false, fmt.Errorf("failed to read file: %w", err)
	}

	binaryCount := 0
	for _, b := range buffer[:n] {
		// ASCII printable range: 9, 10, 13 (tab, newline, carriage return), and 32–126
		if (b > 0 && b < 8) || (b > 13 && b < 32) || b == 0x00 {
			binaryCount++
		}
	}

	if n > 0 && float64(binaryCount)/float64(n) > 0.1 {
		return true, nil
	}

	return false, nil
}
