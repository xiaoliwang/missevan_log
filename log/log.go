package log

import (
    "fmt"
    "time"
    "strings"
)

type Log struct {
    Ip string
    Refer string
    Time time.Time
}

func NewLog(s string) *Log {
    fileds := nginxFileds(s)
    length := len(fileds)
    if 11 == length || 10 == length {
        var refer string
        ip := strings.Trim(fileds[9], "\"")
        ip = parseIp(ip)
        time_str := fileds[3]
        then := parseTime(time_str)
        if 11 == length {
            refer = strings.Trim(fileds[10], "\"")
        } else {
            refer = strings.Trim(fileds[7], "\"")
        }
        
        return &Log{ ip, refer, then }
    }
    return nil
}

func nginxFileds(s string) []string {
    bracket := 0
    quotation := false

    fileds := strings.FieldsFunc(s, func(r rune) bool {
        switch (r) {
            case '[':
                bracket++
            case ']':
                bracket--
            case '"':
                quotation = !quotation
        }
        return !quotation && 0 == bracket && ' ' == r
    });

    return fileds
}

const layout = "2006-Jan-02T15:04:05 (MST)"
func parseTime(s string) time.Time {
    s = s[1:len(s)-1]
    tFileds := strings.FieldsFunc(s, func(r rune) bool {
        return '/' == r || ':' == r || ' ' == r
    })

    s = fmt.Sprintf("%s-%s-%sT%s:%s:%s (CST)", tFileds[2], tFileds[1],  
        tFileds[0], tFileds[3], tFileds[4], tFileds[5]);
    t, err := time.Parse(layout, s)
    if err != nil {
        fmt.Println(err)
    }
    return t
}

func parseIp(s string) string {
    s = strings.Trim(s, "\"")
    var ip string
    for n, r := range s {
        if (',' == r) {
            ip = s[:n]
            break;
        }
    }
    if ("" == ip) {
        ip = s
    }
    return ip
}