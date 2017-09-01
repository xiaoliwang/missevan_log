package mylog

import (
    "testing"
    "fmt"
)

func TestNew(t *testing.T) {
    line := `100.116.223.75 - - [28/Aug/2017:03:18:01 +0800] "GET /mobile/site/getChannelByCatalog?&token=rIPOm1mcZ22bcZaZj1OfmM%2BcWGyEjatommNjkqtmbJxjvdptmZaWWZJa2NjIo5%2BYz5xYbISNq2iaY2OSq2ZsnGO92m2ZlpZZklrI0sSanVmcpauezl1YmpuhoJ%2BaUm9oXYPOmqbQqKnSWp2Hy6Wlp5yTZY6RpKqSrJuVZKOZqKuW18alZcWipMJnxNvEpZKp1ZNlm8WgpGFpYKKknVJhWprE1KWa0Z%2Bm2Fqdh4ZqYpqSnJqfhZSalmmXlqNZZ2VxlMSepFqXY26fm5zShmhhcMWab1SOU6qgo5egWHBSanCUkptnaJNsZ8malsaal2Wcl2iZZ5lqsmJsamtmbGhlbGrdl51pw5ecmZmbypiWk2iXa1hehJauoaGkl1hwYWlxYpiZbWuWbGOIrMzSyFNraJZvb2KYaWZlca8%3D&cid=67 HTTP/1.0" 200 5221 "-" "MissEvanApp/3.5.1 (Android;7.0;HWEVA)" "27.18.246.242" "mi`
    log := NewLog(line)
    fmt.Print(log)
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
