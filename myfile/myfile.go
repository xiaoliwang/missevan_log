package myfile

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"log"
	"time"
	"path"
)

// 复制文件
func Copy(dis, src string)(err error) {
	reader, err := os.Open(src)
	if nil != err {
		return
	}
	defer reader.Close()

	writer, err := CreateFile(dis)
	if nil != err {
		return
	}
	defer func() {
		cerr := writer.Close()
		if nil == err {
			err = cerr
		}
	}()

	var bytes int64
	if bytes, err = io.Copy(writer, reader); nil != err {
		log.Println(err)
		return
	}
	log.Println(src, bytes)

	err = writer.Sync()
	return
}

// 解压文件
func Ungzip(dis, src string) {
	if !FileExist(src) {
		fmt.Println("file unzip not exists", src)
		os.Exit(1)
	}

	defer os.Remove(src);

	zip_file, err := os.Open(src)
	if nil != err {
		fmt.Println(err)
	}
	defer zip_file.Close()

	reader, err := gzip.NewReader(zip_file)
	if nil != err {
		fmt.Println(err)
	}
	defer reader.Close()

	unzip_file, err := CreateFile(dis)
	if nil != err {
		fmt.Println(err)
	}
	defer unzip_file.Close()

	writer := bufio.NewWriter(unzip_file)
	if _, err := io.Copy(writer, reader); nil != err {
		log.Fatal("unzip gz file failed", src, err)
	}
	writer.Flush()
}

func CreateFile(file_path string) (*os.File, error) {
	dir_path := path.Dir(file_path)
	if !DirExist(dir_path) {
		os.MkdirAll(dir_path, 755)
	}
	file, err := os.Create(file_path)
	return file, err
}

// 文件是否存在
func FileExist(file_path string) bool {
	info, err := os.Stat(file_path)
	if nil == err && !info.IsDir() {
		return true
	}
	return false
}

// 文件夹是否存在
func DirExist(file_path string) bool {
	info, err := os.Stat(file_path)
	if nil == err && info.IsDir() {
		return true
	}
	return false
}

// 获取日志文件全路径
const file_name = "20060102"
const dir_name = "200601"
func GetFileNames(date string) []string {
	var today time.Time
	var err error
	if today, err = time.Parse(file_name, date); nil != err {
		log.Fatal("时间格式不合法", err)
		panic(err)
	}

	yesterday := today.AddDate(0, 0, -1)
	tomorrow := today.AddDate(0, 0, 1)

	dates := [3]time.Time {
		tomorrow,
		today,
		yesterday,
	}

	file_type := 4
	file_names := make([]string, len(dates) * file_type)
	// 201708/20170830-1.gz
	for i, date := range dates {
		for j := 1; j <= file_type; j++ {
			filename := fmt.Sprintf("%s/%s-%d", date.Format(dir_name), date.Format(file_name), j)
			file_names[i * file_type + j - 1] = filename
		}
	}
	return file_names
}