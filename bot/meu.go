package bot

import (
	"fmt"
	"github.com/marcmak/calc/calc"
	"github.com/nlopes/slack"
	"log"
	"regexp"
)

type Meu struct {
	*BaseBot
}

var (
	calc_re *regexp.Regexp
)

func NewMeu(token string, stop *chan struct{}) *Meu {
	calc_re = regexp.MustCompile("^계산하라 메우 (.+)")
	return &Meu{NewBot(token, stop)}
}

func (bot *Meu) onMessageEvent(e *slack.MessageEvent) {
	switch {
	case e.Text == "메우, 멱살":
		bot.sendSimple(e, "사람은 일을 하고 살아야한다. 메우")
		break
	case e.Text == "메우메우 펫탄탄":
		bot.sendSimple(e, `메메메 메메메메 메우메우
메메메 메우메우
펫땅펫땅펫땅펫땅 다이스키`)
		break
	default:
		if matched, ok := MatchRE(e.Text, calc_re); ok {
			defer func() {
				if r := recover(); r != nil {
					log.Println("Recovered : %g", r)
					bot.replySimple(e, "에러났다 메우")
				}
			}()
			bot.sendSimple(e, fmt.Sprintf("%f", calc.Solve(matched[1])))
		}
	}
}
