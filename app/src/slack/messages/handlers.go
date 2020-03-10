package messages

import (
	"math/rand"
	"strings"
	"time"

	"golang.org/x/exp/utf8string"

	"slack/todo"
)

var hello = func() string {
	return "やあ！:raised_hand:\nぼくはホリネズミのGopher(ゴーファー)だよ。"
}

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
	task = strings.TrimSpace(task)
	return
}

var mouse = func() string { return "ちがう！:rage:" }

var gophers = func() string { return "そうだよ！:smile:" }

func nilTaskName() string { return "タスクが未入力だよ:sweat:" }

func notTask(task string) string {
	msg := []string{task, " っていうタスクが存在しないよ:disappointed_relieved:"}
	return strings.Join(msg, "")
}

// Todoの追加
var todoAdd = func() string {
	t := getTask()
	if t == "" { // タスクの未入力時
		return nilTaskName()
	}

	todo.Add(t)
	m := []string{t, " を追加したよ！"}
	return strings.Join(m, "")
}

// Todoの完了
var todoDone = func() string {
	t := getTask()
	if t == "" { // タスクの未入力時
		return nilTaskName()
	}

	if !todo.Done(t) {
		// タスクが存在しなかった時
		return notTask(t)
	}

	m := []string{t, " を完了にしたよ！"}
	return strings.Join(m, "")
}

// Todoの削除
var todoDel = func() string {
	t := getTask()
	if t == "" { // タスクの未入力時
		return nilTaskName()
	}

	if !todo.Del(t) {
		// タスクが存在しなかった時
		return notTask(t)
	}
	m := []string{t, " を削除したよ！"}
	return strings.Join(m, "")
}

// 未完了Todo一覧
var todoList = func() string {
	t, err := todo.List()
	if err != nil {
		return "未完了のタスクは今はないよ！"
	}

	msg := make([]string, len(t)+1)
	msg[0] = "これが未完了のタスクだよ！"
	for i := 0; i < len(t); i++ {
		msg[i+1] = t[i]
	}
	return strings.Join(msg, "\n")
}

// 完了Todo一覧
var todoDoneList = func() string {
	t, err := todo.DoneList()
	if err != nil {
		return "完了済みのタスクは今はないよ！"
	}

	msg := make([]string, len(t)+1)
	msg[0] = "これが完了済みのタスクだよ！"
	for i := 0; i < len(t); i++ {
		msg[i+1] = t[i]
	}
	return strings.Join(msg, "\n")
}
