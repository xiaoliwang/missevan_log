package myfile

import (
	"testing"
	"fmt"
)

//func TestCopy(t *testing.T) {
//	if err := Copy("./test.file", "/oss/datas/201708/20170830-3.gz"); nil != err {
//		fmt.Println(err)
//	}
//}

func TestFilename(t *testing.T) {
	file_names := GetFileNames("20170831")
	fmt.Println(file_names)
}