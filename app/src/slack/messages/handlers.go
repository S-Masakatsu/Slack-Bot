package messages

import (
	"math/rand"
	"strings"
	"time"

	"golang.org/x/exp/utf8string"

	"slack/todo"
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

// 送信されたメッセージを整形し、タスクの部分だけ切り取り、
// タスク名を返します。
func getTask() (task string) {
	m := utf8string.NewString(getMessage)
	si := strings.Index(getMessage, ":")
	task = m.Slice(si+1, m.RuneCount())
	return
}

var mouse = func() string { return "ちがう！" }

var gophers = func() string { return "そうだよ！" }

// Todoの追加
var todoAdd = func() string {
	t := getTask()
	todo.Add(t)
	m := []string{"ToDo: ", t, " を追加したよ！"}
	return strings.Join(m, "")
}

// Todoの完了
var todoDone = func() string {
	t := getTask()
	todo.Done(t)
	m := []string{"ToDo: ", t, " を完了にしたよ！"}
	return strings.Join(m, "")
}

// Todoの削除
var todoDel = func() string {
	t := getTask()
	todo.Del(t)
	m := []string{"ToDo: ", t, " を削除したよ！"}
	return strings.Join(m, "")
}

// 未完了Todo一覧
var todoList = func() string {
	t := todo.List()
	msg := make([]string, len(t)+1)
	msg[0] = "これが未完了のタスクだよ！"
	for i := 0; i < len(t); i++ {
		msg[i+1] = t[i]
	}
	return strings.Join(msg, "\n")
}

// 完了Todo一覧
var todoDoneList = func() string {
	t := todo.DoneList()
	msg := make([]string, len(t)+1)
	msg[0] = "これが完了済みのタスクだよ！"
	for i := 0; i < len(t); i++ {
		msg[i+1] = t[i]
	}
	return strings.Join(msg, "\n")
}
