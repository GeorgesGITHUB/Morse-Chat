package main

import (
	"regexp"
	"strings"
	"github.com/gSpera/morse"
	"time"
)

type Message struct {
	Message_id 	 int	   `json:"message_id,string"`
	Sender_id    int       `json:"sender_id,string"`
	ContentRaw   string    `json:"contentRaw"`
	ContentText  string    `json:"contentText"`
	ContentMorse string    `json:"contentMorse"`
	Timestamp    time.Time `json:"timestamp"`
}

func toContentText(contentRaw string) string {
	if isMorse(contentRaw){
		return strings.ToLower(morse.ToText(contentRaw))
	} else {
		return contentRaw
	}
}

func toContentMorse(contentRaw string) string {
	if isMorse(contentRaw){
		return contentRaw
	} else {
		return morse.ToMorse(contentRaw)
	}
}

func isMorse(str string) bool {
	re := regexp.MustCompile(`^[.\-\s]+$`)
	return re.MatchString(str)
}
