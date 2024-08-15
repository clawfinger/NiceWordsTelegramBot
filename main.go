package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"nicebot/internal/config"
	"os"
	"os/signal"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
	"golang.org/x/exp/rand"
)

func main() {
	config.Parse()

	bot, err := tgbotapi.NewBotAPI(config.Config.Token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	ctx, cancFunc := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancFunc()

	config.Config.CronSchedule = "0 10 * * *"

	lastIdx := 0
	scheduler := cron.New()
	_, err = scheduler.AddFunc(config.Config.CronSchedule, func() {
		words, err := getWordsFromPastebin()
		if err != nil {
			log.Println(err.Error())
		}

		rndIdx := rand.Intn(len(words) - 1)
		for rndIdx == lastIdx {
			rndIdx = rand.Intn(len(words) - 1)
		}

		msg := tgbotapi.NewMessage(config.Config.ChannelID, words[rndIdx])
		lastIdx = rndIdx
		bot.Send(msg)
	})
	if err != nil {
		log.Println(err)
	}
	scheduler.Start()
	<-ctx.Done()
	scheduler.Stop()
}

func getWordsFromPastebin() ([]string, error) {
	resp, err := http.Get(config.Config.PastebinUrl)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var words []string
	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Panic(err)
	}
	return words, nil
}
