package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"slack/config"
	"slack/messages"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

func EventPoint(w http.ResponseWriter, r *http.Request, s config.SlsckItems) {
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
			api.PostMessage(ev.Channel, slack.MsgOptionText(messages.PostMessage(txt), false))
		}
	}
}
