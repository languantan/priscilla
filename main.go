package main

import (
	"log"
	"strings"
	"unicode"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("394586798:AAGtPSoVTZbuBYzs4Bhp2xJ7kpXRF6a8hNo")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	punMap := map[string]string{
		"axe":        "I used to work as a lumberjack, but then I got axed.",
		"bank":       "I used to work in a bank, but then I lost interest.",
		"factory":    "I used to work at a blanket factory, but then it folded.",
		"book":       "Books about gravity are hard to put down.",
		"car":        "If you stand behind a car, you'll become exhausted.",
		"coffee":     "I do some of my best thinking over coffee. I tend to have a latte on my mind.",
		"finger":     "Jill broke her finger, but on the other hand she's fine.",
		"flower":     "If a florist shop catches fire, would that be a florist fire?",
		"iphone":     "I named my iphone Titanic, now every time i plug it into my PC, Titanic is syncing.",
		"monorail":   "Puns about monorail always make for decent one-liners.",
		"mug":        "If someone steals my mug, is that considered a mugging?",
		"pasta":      "Did you hear about the Italian chef? he pasta way.",
		"pay":        "I'm so poor I can't even pay attention.",
		"punch":      "When Peter Pan punches, they neverland.",
		"promotion":  "Why did the scarecrow get promoted? Because it was outstanding in its field.",
		"ransomware": "How did the people behind the latest hack escape? they ran-som-ware.",
		"rest":       "The police were called up to a kindergarten because a child was resisting a rest.",
		"seconds":    "When a clock is hungry it goes back four seconds.",
		"tea":        "Do you know what kind of tea is sometimes hard to swallow? Reali-tea.",
		"tennis":     "I had a game of quiet tennis today, it's just like regular tennis but without the racket.",
		"toast":      "My friend used to own a bakery, but his business went toast.",
		"time":       "I made a belt with watches, it was a waist of time.",
		"tired":      "If you stand in front of a car, you'll become tired.",
		"train":      "A crazy person who steals trains must have some locomotives.",
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		userMsg := strings.ToLower(update.Message.Text)

		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}

		words := strings.FieldsFunc(userMsg, f)

		for _, v := range words {
			log.Printf(v)
			if punMapVal, ok := punMap[v]; ok {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, punMapVal)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				break
			}
		}
	}
}
