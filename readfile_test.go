// Author stone-bird created on 2021/8/8 10:04.
// Email 1245863260@qq.com or g1245863260@gmail.com.
// Use of example read line

package readfile

import (
	"fmt"
	"log"
)

func ExampleNewReader() {
	fileName := "test.txt"
	fr ,err := NewReader(fileName)
	if err != nil{
		log.Fatal(err)
	}
	go fr.Scanner()
	for line := range fr.ContentChan{
		fmt.Println(line)
	}

	// Output:
	// line 1
	// line 2
	// line 3
	// ....
}