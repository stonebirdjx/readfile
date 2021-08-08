// Author stone-bird created on 2021/8/8 10:04.
// Email 1245863260@qq.com or g1245863260@gmail.com.
// Use of more goroutine deal the file line content
package readfile

import (
	"bufio"
	"os"
)

// file reader struct
type FileReader struct {
	FileName    string
	ContentChan chan string
	File        *os.File
}

// make a new file reader
func NewReader(fileName string) (*FileReader, error) {
	var fr FileReader
	fr.FileName = fileName
	fr.ContentChan = make(chan string)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fr.File = file
	return &fr, nil
}

//read line recommend scanner
func (fr *FileReader) Scanner() {
	bs := bufio.NewScanner(fr.File)
	for bs.Scan() {
		lineString := bs.Text()
		fr.ContentChan <- lineString
	}
	close(fr.ContentChan)
}
