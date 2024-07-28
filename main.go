package main

import (
	"context"
	"encoding/json"
	"log"
	"nicebot/internal/config"
	"os"
	"os/signal"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/exp/rand"
)

func main() {
	config.Parse()

	bot, err := tgbotapi.NewBotAPI(config.Config.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	chat, err := bot.GetChat(tgbotapi.ChatInfoConfig{
		ChatConfig: tgbotapi.ChatConfig{
			ChatID: config.Config.ChannelID,
		},
	})
	if err != nil {
		log.Panic(err)
	}
	log.Printf("chat name %s", chat.UserName)

	var words []string
	jsonRaw, err := os.ReadFile("nicewords.json")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(jsonRaw, &words)
	if err != nil {
		log.Panic(err)
	}
	ticker := time.NewTicker(time.Duration(config.Config.Timeout) * time.Minute)

	ctx, cancFunc := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancFunc()
	for {
		select {
		case <-ticker.C:
			log.Print("starting send cycle")
			rndIdx := rand.Intn(len(words) - 1)
			msg := tgbotapi.NewMessage(config.Config.ChannelID, words[rndIdx])
			bot.Send(msg)
		case <-ctx.Done():
			log.Print("shutting down")
			os.Exit(0)
		}
	}
}
