package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

type BotNmae struct {
	Members []struct {
		Name string `json:"name"`
	}
}

type Slscks struct {
	Token   string
	Vtoken  string
	BotNmae string
}

type KeyWard struct {
	Key    string
	Return func() string
}

type KeyWards []KeyWard

func init() {
	// 環境変数：BOT_UNAMEを設定
	envBotName()

	s := Slscks{
		Token:   readToken(os.Getenv("BOT_USER_OAUTH_ACCESS_TOKEN_PATH")),
		Vtoken:  readToken(os.Getenv("VERIFICATION_TOKEN_PATH")),
		BotNmae: os.Getenv("BOT_UNAME"),
	}
	http.HandleFunc("/v1/event-point", s.eventPoint)
}

// users.list APIからBot nameを取得し、環境変数を生成する
func envBotName() {
	par := url.Values{}
	t := readToken(os.Getenv("WORKSPACE_TOKEN_PATH"))
	par.Add("token", t)
	uri := []string{os.Getenv("SLACK_USER_LIST_API"), "?", par.Encode()}
	r, err := http.Get(strings.Join(uri, ""))
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	bName := new(BotNmae)
	if err := json.Unmarshal([]byte(body), bName); err != nil {
		panic(err)
	}
	os.Setenv("BOT_UNAME", bName.Members[0].Name)
}

// ファイルからTokenを取得する
func readToken(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (s *Slscks) eventPoint(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	eventsAPI, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: s.Vtoken}))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if eventsAPI.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}

	api := slack.New(s.Token)
	if eventsAPI.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPI.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent: // Botユーザーへのメンションの場合
			log.Println("AppMentionEvent")
			txt := ev.Text
			api.PostMessage(ev.Channel, slack.MsgOptionText(postMessage(txt), false))
		}
	}
}

func postMessage(text string) (msg string) {
	for _, k := range keyWord {
		if strings.Contains(text, k.Key) {
			msg = k.Return()
		}
	}
	return
}

var keyWord = KeyWards{
	KeyWard{
		"やあ",
		hello,
	},
	KeyWard{
		"おみくじ",
		returnFortune,
	},
}

func hello() string {
	return "やあ。ぼくはホリネズミのGopherだよ。"
}

func returnFortune() string {
	f := []string{"大吉", "吉", "中吉", "末吉", "凶"}
	rand.Seed(time.Now().UnixNano())
	msg := []string{"今日は、", f[rand.Intn(len(f)-1)], "だよ！"}
	return strings.Join(msg, "")
}

func main() {
	port := []string{":", os.Getenv("APP_PORT")}
	fmt.Println("[INFO] Server listening")
	if err := http.ListenAndServe(strings.Join(port, ""), nil); err != nil {
		log.Panicln(err)
		os.Exit(1)
	}
}
