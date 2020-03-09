package todo

// Slack Bot がToDoを管理します。
// ----- Commands ------
// @BotName todo				-> Todo を追加します。
// @BotName done 				-> Todo を完了にします。
// @BotName del 				-> Todo を削除します。
// @BotName list 				-> 未完了の Todo を一覧表示します。
// @BotName donelist 		-> 完了済みの Todo を一覧表示します。

import "strings"

type Todo struct {
	Todo string // Task name
	Done bool   // Boolean value for Done
}

type TodoList []Todo // Structure to store Todo

var Todos TodoList // Array for saving Todo

// 文字列から新たにタスクを作成し、リストに追加します。
func Add(task string) {
	var t Todo
	t.Todo = strings.TrimSpace(task)
	Todos = append(Todos, t)
}

// タスクを受け取り、タスクが完了したかどうかを返します。
// 完了済みなら => true / 未完了なら => false true を返します。
func isDone(t Todo) bool {
	return t.Done
}

// タスクを受け取り、タスクが完了していないかどうかを返します。
// 完了済みなら => false / 未完了なら => true を返します。
func isNotDone(t Todo) bool {
	return !isDone(t)
}

// 未完了のタスクをリスト化して返します。
func List() (t []string) {
	for _, key := range Todos {
		if isNotDone(key) {
			t = append(t, key.Todo)
		}
	}
	return
}

// タスクを完了状態にします。
func Done(task string) {
	for i, key := range Todos {
		if key.Todo == task {
			Todos[i].Done = true
		}
	}
}

// 完了済みのタスクをリスト化して返します。
func DoneList() (t []string) {
	for _, key := range Todos {
		if isDone(key) {
			t = append(t, key.Todo)
		}
	}
	return
}

// タスクを削除します。
func Del(task string) {
	for i, k := range Todos {
		if k.Todo == strings.TrimSpace(task) {
			Todos = append(Todos[:i], Todos[i+1:]...)
		}
	}
}
