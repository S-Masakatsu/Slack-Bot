package messages

import (
	"math/rand"
	"strings"
	"time"
)

var hello = func() string { return "やあ！\nぼくはホリネズミのGopher(ゴーファー)だよ。" }

var fortune = func() string {
	f := []string{"大吉", "吉", "中吉", "末吉", "凶"}
	rand.Seed(time.Now().UnixNano())
	msg := []string{"今日の運勢は、「", f[rand.Intn(len(f)-1)], "」だよ！"}
	return strings.Join(msg, "")
}

var document = func() string {
	msg := `公式ドキュメント：https://golang.org/
GoDoc：https://godoc.org/
APIドキュメント：https://go.dev/
より良いコードを書くため：https://github.com/golang/go/wiki/CodeReviewComments`
	return msg
}

var tutorial = func() string {
	msg := `これがオススメだよ！
A Tour of Go：https://go-tour-jp.appspot.com/list`
	return msg
}

var mouse = func() string { return "ちがう！" }

var gophers = func() string { return "そうだよ！" }
