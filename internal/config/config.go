package config

import (
	"github.com/caarlos0/env/v11"
)

type config struct {
	Token        string `env:"BOT_TOKEN"`
	ChannelID    int64  `env:"CHANNEL_ID"`
	Timeout      int64  `env:"MSG_TIMEOUT_MINS"`
	PastebinUrl  string `env:"PASTEBIN_URL"`
	CronSchedule string `env:"CRON_SCHEDULE"`
}

var Config config

func Parse() error {
	return env.Parse(&Config)
}
