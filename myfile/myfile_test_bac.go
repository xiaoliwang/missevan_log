package myfile

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var from = "./gz/"
var dist = "./dist/"

func clearAll() {
	if (DirExist(dist)) {
		os.RemoveAll(dist);
	}
}

func TestMain(m *testing.M) {
	clearAll()
	ret := m.Run()
	clearAll()
	os.Exit(ret)
}

func TestCopy(t *testing.T) {
	if !DirExist(dist) {
		os.Mkdir(dist, 755)
	}
	files, err := ioutil.ReadDir(from)
	if nil != err {
		t.Error(err)
	}
	for _, file := range files {
		Copy(dist+file.Name(), from+file.Name())
	}
}

func TestUngzip(t *testing.T) {
	files, err := ioutil.ReadDir(dist)
	if nil != err {
		t.Error(err)
	}

	for _, file := range files {
		file_name := file.Name()
		new_name := strings.Replace(file_name, ".gz", "", 1)
		Ungzip(dist+new_name, dist+file_name)

	}
}
