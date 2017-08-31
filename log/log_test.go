package log

import (
    "testing"
    "fmt"
)

func TestNew(t *testing.T) {
    line := `100.116.223.59 - - [30/Aug/2017:03:49:01 +0800] "POST /mobile/site/addplaytimes HTTP/1.0" 200 77 "-" "MissEvanApp/3.5.2 (Android;4.4.4;msm8916_32)" "117.141.67.237" "missevan"`
    log := NewLog(line)
}

/*func TestNew(t *testing.T) {
    var (
        line string
        part []byte
        prefix bool
    )

    defer func() {
        if err := recover(); err != nil {
            t.Error(err, line)
        }
    }()

    path := "./test_file/20170727-2"
    r_file, err := os.Open(path)
    if nil != err {
        t.Error("Open file failed", err)
    }
    defer r_file.Close()
    reader := bufio.NewReader(r_file)
    buffer := bytes.NewBuffer(make([]byte, 1024))

    for {
        if part, prefix, err = reader.ReadLine(); nil != err {
            break
        }
        buffer.Write(part)
        if !prefix {
            line = buffer.String()
            buffer.Reset()
            if log := NewLog(line); nil == log {
                continue
            }
        }
    }

    t.Log("SUCCESS")
}*/
