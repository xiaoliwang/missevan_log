package mylog

import (
	"time"
	"strings"
	"fmt"
	_ "log"
	"log"
)

type Log struct {
	Ip string
	From string
	Time time.Time
}

func NewLog(s string) (my_log *Log) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			my_log = nil
		}
	}()
	fields := parseNginxLog(s)
	length := len(fields)
	if 11 == length || 10 == length {
		my_log = &Log{
			ipFormat(fields[9]),
			fromFormat(fields, length),
			timeFormat(fields[3]),
		}
	}
	return
}

func parseNginxLog(s string) []string {
	bracket := 0
	quotation := false

	fields := strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case '[':
			bracket++
		case ']':
			bracket--
		case '"':
			quotation = !quotation
		}
		return !quotation && 0 == bracket && ' ' == r
	})

	return fields
}

func ipFormat(s string) string {
	s = strings.Trim(s, "\"")
	var ip string
	for n, r := range s {
		if ',' == r {
			ip = s[:n]
			break;
		}
	}
	if "" == ip {
		ip = s
	}
	return ip
}

const layout = "2006-Jan-02T15:04:05 (MST)"
func timeFormat(s string) time.Time {
	s = s[1:len(s) - 1] // 去除中括号
	timeFields := strings.FieldsFunc(s, func(r rune) bool {
		return '/' == r || ':' == r || ' ' == r
	})

	s = fmt.Sprintf("%s-%s-%sT%s:%s:%s (CST)", timeFields[2], timeFields[1],
		timeFields[0], timeFields[3], timeFields[4], timeFields[5]);
	t, err := time.Parse(layout, s)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

var froms = [...]string {
	// "missevan",
	"missevan_google",
	"missevan_baidu",
	"missevan_360",
	"missevan_mi",
	"missevan_oppo",
	"missevan_vivo",
	"missevan_huawei",
	"missevan_tencent",
	"missevan_nubia",
	"missevan_gionee",
	"missevan_wandoujia",
}

func fromFormat(fields []string, length int) string {
	var from string

	if 11 == length {
		from = strings.Trim(fields[10], "\"")
	} else {
		from = strings.Trim(fields[7], "\"")
	}

	for _, new_from := range froms {
		if from == new_from {
			return from
		}
	}
	return "-"
}

