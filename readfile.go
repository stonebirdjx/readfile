package readfile

import (
	"bufio"
	"os"
)

/*
* @Auther:create by hjx
* @Email :1245863260@qq.com g1245863260@gmail.com
* @Date  :2021/8/8 10:15
 */

// file reader struct
type FileReader struct {
	FileName    string
	ContentChan chan string
	File        *os.File
}

// make a new file reader
func NewFileReader(fileName string) (*FileReader, error) {
	var fileReader FileReader
	fileReader.FileName = fileName
	fileReader.ContentChan = make(chan string)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fileReader.File = file
	return &fileReader, nil
}

// read line recommend scanner
func (fileReader *FileReader) BufioScanner() {
	bs := bufio.NewScanner(fileReader.File)
	for bs.Scan(){
		lineString := bs.Text()
		fileReader.ContentChan <- lineString
	}
	close(fileReader.ContentChan)
}