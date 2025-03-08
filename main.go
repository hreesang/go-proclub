package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hreesang/go-proclub/bot"
	_ "github.com/hreesang/go-proclub/bot/config"
	_ "github.com/hreesang/go-proclub/bot/database"
	_ "github.com/hreesang/go-proclub/bot/database/models"
	_ "github.com/hreesang/go-proclub/bot/proclubs"
	_ "github.com/hreesang/go-proclub/bot/slashcommands"
	_ "github.com/hreesang/go-proclub/bot/utils"
)

func main() {
	if err := bot.Run(); err != nil {
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<- stop
	bot.Stop()
}
