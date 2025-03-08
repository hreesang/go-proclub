package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/hreesang/go-proclub/bot/utils"
)

var (
	BotToken		string
	MongoDBUri		string
	MongoDBDatabase	string
)

func init() {
	if err := godotenv.Load(); err != nil {
		utils.Log.Fatal("failed to load environment variables")
		return
	}

	BotToken = strings.TrimSpace(os.Getenv("BOT_TOKEN"))
	if BotToken == "" {
		utils.Log.Fatal("failed to load 'BOT_TOKEN' from the environment variables")
		return
	}

	MongoDBUri = strings.TrimSpace(os.Getenv("MONGODB_URI"))
	if BotToken == "" {
		utils.Log.Fatal("failed to load 'MONGODB_URI' from the environment variables")
		return
	}

	MongoDBDatabase = strings.TrimSpace(os.Getenv("MONGODB_DATABASE"))
	if BotToken == "" {
		utils.Log.Fatal("failed to load 'MONGODB_DATABASE' from the environment variables")
		return
	}
}
