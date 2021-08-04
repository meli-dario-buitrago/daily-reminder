package main

import (
	"os"
	"os/signal"

	"github.com/meli-dario-buitrago/daily-reminder/config"
	"github.com/meli-dario-buitrago/daily-reminder/slack"
	"github.com/robfig/cron"
)

func main() {
	config.ConfigureProperties()
	c := cron.New()
	c.AddFunc("0 25 9 * * *", slack.SendSlackMessage)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
