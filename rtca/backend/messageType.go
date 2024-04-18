package main

import (
	"github.com/gSpera/morse"
	"regexp"
	"strings"
)

type Message struct { // Custom type for messages
	Sender       string `json:"sender"`
	ContentRaw   string `json:"contentRaw"`
	ContentText  string `json:"contentText"`
	ContentMorse string `json:"contentMorse"`
}

func (msg *Message) fillMissingUsingRaw() {
	if isMorse(msg.ContentRaw) {
		msg.ContentText = strings.ToLower(morse.ToText(msg.ContentRaw))
		msg.ContentMorse = msg.ContentRaw
	} else {
		msg.ContentMorse = morse.ToMorse(msg.ContentRaw)
		msg.ContentText = msg.ContentRaw
	}
}

func isMorse(str string) bool {
	re := regexp.MustCompile(`^[.\-\s]+$`)
	return re.MatchString(str)
}
