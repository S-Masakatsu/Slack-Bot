package messages

import "strings"

func PostMessage(text string) (msg string) {
	for _, k := range actions {
		if strings.Contains(text, k.KeyMessage) {
			getMessage = text
			msg = k.HandleFunc()
		}
	}
	return
}
