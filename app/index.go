package main

import (
	"flag"
	"log"
	"os"
	"missevan_log/myfile"
	"bufio"
	"bytes"
	"missevan_log/mylog"
	"fmt"
	"sync"
	"time"
	"path"
)

var rwg sync.WaitGroup
var wwg sync.WaitGroup

var date string
var input string
var output string

var temp string
var download bool

var chan_log = make(chan *mylog.Log, 30)

func init() {
	flag.StringVar(&date, "date", "", "the date you want to get")
	flag.StringVar(&input, "input", "./input/", "where you want to get the files")
	flag.StringVar(&output, "output", "./output/", "where you put the files")
	flag.BoolVar(&download, "download", true, "do you want to download files")
	flag.Parse()
}

func main() {
	// 记录报错信息
	defer func() {
		if err := recover(); nil != err {
			log.Fatal(err)
		}
	}()

	if !myfile.DirExist(output) {
		os.Mkdir(output, 755)
	}

	temp = path.Join(output, "temp")
	if !myfile.DirExist(temp) {
		os.Mkdir(temp, 755)
	}

	getLogs(date)
	go writeLog()
	go alive()
	rwg.Wait()
	close(chan_log)
	wwg.Wait()
}

func alive() {
	for {
		fmt.Printf("%c", '.')
		time.Sleep(5 * time.Second)
	}
}

func writeLog() {
	map_chans := make(map[string]chan *mylog.Log)
	for my_log := range chan_log {
		file_name := path.Join(output, date, "files", my_log.From + ".txt")
		if !myfile.FileExist(file_name) {
			file, err := myfile.CreateFile(file_name)
			if nil != err {
				log.Fatal(err)
			}

			my_chan := make(chan *mylog.Log)
			map_chans[my_log.From] = my_chan
			defer close(my_chan)

			wwg.Add(1)
			go func(file *os.File, my_chan chan*mylog.Log) {
				defer wwg.Done()
				defer file.Close()
				writer := bufio.NewWriter(file)
				for my_log := range my_chan {
					fmt.Fprintln(writer, my_log.Ip)
				}
				writer.Flush()
			}(file, my_chan)
		}
		map_chans[my_log.From] <- my_log
	}
}

func getLogs(date string) {
	file_names := myfile.GetFileNames(date)
	for _, file_name := range file_names {
		rwg.Add(1)
		go downloadFiles(file_name)
	}
}

func downloadFiles(file_name string) {
	defer rwg.Done()
	gz_file := file_name + ".gz"
	initial_file := path.Join(input, gz_file)
	temp_file := path.Join(temp, gz_file)
	unzip_file := path.Join(temp, file_name)

	if download || !myfile.FileExist(unzip_file) {
		myfile.Copy(temp_file, initial_file)
		myfile.Ungzip(unzip_file, temp_file)
	}

	handleLog(unzip_file)
}

func handleLog(file_name string) {
	var (
		line string
		part []byte
		prefix bool
	)

	defer func() {
		if err := recover(); err != nil {
			log.Println(line)
			log.Fatal(err)
		}
	}()

	r_file, err := os.Open(file_name)
	if nil != err {
		log.Fatal("Open file failed", err)
	}
	defer r_file.Close()

	reader := bufio.NewReader(r_file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); nil != err {
			break
		}
		buffer.Write(part)
		if !prefix {
			line = buffer.String()
			if my_log := mylog.NewLog(line); nil != my_log && "-" != my_log.From {
				if date == my_log.Time.Format("20060102") {
					chan_log <- my_log
				}
			}
			buffer.Reset()
		}
	}
}