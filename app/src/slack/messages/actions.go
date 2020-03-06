package messages

type Action struct {
	KeyMessage string        // キーとなるメッセージ
	HandleFunc func() string // 返すメッセージの関数
}

type Actions []Action

var actions = Actions{
	Action{
		"やあ",
		hello,
	},
	Action{
		"おみくじ",
		fortune,
	},
	Action{
		"ドキュメント",
		document,
	},
	Action{
		"チュートリアル",
		tutorial,
	},
	Action{
		"ねずみ",
		mouse,
	},
	Action{
		"ホリネズミ",
		gophers,
	},
}
