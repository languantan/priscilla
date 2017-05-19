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
		"axe":          "I used to work as a lumberjack, but then I got axed.",
		"baker":        "Bakers trade recipes on a knead to know basis.",
		"bank":         "I used to work in a bank, but then I lost interest.",
		"batman":       "What do you get when Batman leaves church early? Christian Bale.",
		"book":         "Books about anti gravity are impossible to put down.",
		"bicycle":      "A bicycle can't stand on its own because it's two tired.",
		"bridge":       "What did the man say when the bridge fell on him? The suspension is killing me.",
		"car":          "If you run behind a car, you'll become exhausted.",
		"calendar":     "I don't like using calendars. It feels like my days are numbered.",
		"coffee":       "I do some of my best thinking over coffee. I tend to have a latte on my mind.",
		"constipation": "Did you hear about the new movie 'Constipation'? It hasn't come out yet.",
		"clock":        "Have you ever tried to eat a clock? It's very time consuming.",
		"stable":       "An injured half-man half-horse creature was brought to the hospital. Doctor says his condition is stable.",
		"england":      "England doesn't have a kidney bank, but it has a liverpool.",
		"factory":      "I used to work at a blanket factory, but then it folded.",
		"food":         "Yesterday I accidentally swallowed some food coloring. The doctor says I'm fine but I feel like I've dyed a little on the inside.",
		"finger":       "Jill hurt her finger, but on the other hand, she's fine.",
		"flower":       "If a flower shop catches fire, would that be a florist fire?",
		"iphone":       "I named my iphone Titanic, now every time i plug it into my PC, Titanic is syncing.",
		"magician":     "The magician got so mad he pulled his hare out.",
		"medium":       "What do you call a midget fortune teller that just escaped from prison? Small medium at large.",
		"monorail":     "Puns about monorail always make for decent one-liners.",
		"mug":          "If someone steals my mug, is that considered a mugging?",
		"pasta":        "Did you hear about the Italian chef? he pasta way.",
		"pay":          "I'm so poor I can't even pay attention.",
		"pencil":       "Pencils could be made with erasers on both ends, but what would be the point?",
		"pigeon":       "Did you hear about the pigeon with diarrhea? It kept saying poopoo poopoo.",
		"punch":        "When Peter Pan punches, they neverland.",
		"promotion":    "Why did the scarecrow get promoted? Because it was outstanding in its field.",
		"hack":         "How did the people behind the latest hack escape? they ransomware.",
		"hungry":       "If a clock gets hungry, it goes back four seconds.",
		"rest":         "The police were called up to a kindergarten because a child was resisting a rest.",
		"sleep":        "Sleeping comes naturally to me, I can do it with my eyes closed.",
		"tea":          "Do you know what kind of tea is sometimes hard to swallow? Reali-tea.",
		"tennis":       "I had a game of quiet tennis today, it's just like regular tennis but without the racket.",
		"toast":        "My friend used to own a bakery, but then his business went toast.",
		"time":         "I made a belt with watches, it was a waist of time.",
		"tired":        "If you run in front of a car, you'll become tired.",
		"train":        "A crazy person who steals trains must have some locomotives.",
		"vacuum":       "Being the owner of a vacuum cleaner company must suck.",
		"watch":        "I was going to look for my missing watch, but I could never find the time.",
		"whiteboard":   "Whiteboards are remarkable.",
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
