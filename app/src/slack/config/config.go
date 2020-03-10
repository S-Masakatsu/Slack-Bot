package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	m "middlewares"
)

const (
	tokenPath  = "BOT_USER_OAUTH_ACCESS_TOKEN_PATH"
	vtokenPath = "VERIFICATION_TOKEN_PATH"
	workPath   = "WORKSPACE_TOKEN_PATH"
	envBot     = "BOT_UNAME"
	api        = "SLACK_USER_LIST_API"
)

// SlackAPI[user/loist]からname を取得した際の保管用の構造体です。
type Names struct {
	Members []struct {
		Name string `json:"name"`
	}
}

// パッケージ nlopes/slack で必要な情報の構造体です。
type SlsckItems struct {
	Token   string
	Vtoken  string
	BotName string
}

func init() {
	// Set environment variable BOT_UNAME
	setBotName()
}

// 指定したパスからファイルを読み取ります。
// 呼び出しに失敗した場合はerrorが発生します。
func getToken(path string) string {
	b, e := ioutil.ReadFile(path)
	if e != nil {
		err := []string{"Call failed: ", e.Error()}
		panic(strings.Join(err, ""))
	}
	return string(b)
}

// BotNameをSlackAPIから取得し環境変数(BOT_UNAME)を設定します。
// APIのリダイレクトが多すぎる場合，またはHTTPプロトコルerror が発生した場合は，errorが発生します。
func setBotName() {
	// Set the parameters.
	par := url.Values{}
	t := getToken(m.GetEnv(workPath))
	par.Add("token", t)
	// Creates a ULI based on the set parameters and sends a request to the API.
	url := []string{m.GetEnv(api), "?", par.Encode()}
	r, err := http.Get(strings.Join(url, ""))
	if err != nil {
		panic(err)
	}

	// Extract only the name from the response (JSON) from the API, and set the structure (BotNmae).
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	bName := new(Names)
	if err := json.Unmarshal([]byte(body), bName); err != nil {
		panic(err)
	}
	// Set environment variables from the structure (BotName).
	os.Setenv(envBot, bName.Members[0].Name)
}

// 構造体 SlsckItems を設定して、返します。
func GetSlackItem() (s SlsckItems) {
	token := getToken(m.GetEnv(tokenPath))
	vtoken := getToken(m.GetEnv(vtokenPath))
	name := m.GetEnv(envBot)
	s = SlsckItems{
		Token:   token,
		Vtoken:  vtoken,
		BotName: name,
	}
	return
}
