package main

import (
	"strings"

	"./user"
	"github.com/turnage/graw/reddit"
)

type liskTipBot struct {
	bot reddit.Bot
}

func (r *liskTipBot) Comment(p *reddit.Comment) error {
	if strings.Contains(p.Body, "!lisktip ") {
		amount := strings.SplitAfter(p.Body, "!lisktip ")[1]
		amount = strings.Split(amount, " ")[0]
		initiateTip(p.Author, amount, p.ID)
		return r.bot.Reply(p.Name, "This triggered lisk tip for an amout of "+amount+" LSK")
	}
	if strings.Contains(p.Body, "!tiplisk ") {
		amount := strings.SplitAfter(p.Body, "!tiplisk ")[1]
		amount = strings.Split(amount, " ")[0]
		initiateTip(p.Author, amount, p.ID)
		return r.bot.Reply(p.Name, "This triggered lisk tip for an amout of "+amount+" LSK")
	}
	return nil
}

func initiateTip(username string, amount string, commentID string) error {
	usr := user.GetUser(username)
	usr.GetUserData()
	usr.SendLisk(amount, username)
	return nil
}
