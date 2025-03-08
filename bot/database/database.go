package database

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/bwmarrin/discordgo"
	"github.com/hreesang/go-proclub/bot"
	"github.com/hreesang/go-proclub/bot/config"
	"github.com/hreesang/go-proclub/bot/utils"
)

var (
	setupOnce 		sync.Once
	setupDatabase 	*mongo.Database
	setupError		error
)

func Setup() (*mongo.Database, error) {
	setupOnce.Do(func () {
		clientOptions := options.Client().
			ApplyURI(config.MongoDBUri)
		
		client, err := mongo.Connect(clientOptions)
		if err != nil {
			setupError = err
			return
		}

		setupDatabase = client.Database(config.MongoDBDatabase)
	})
	
	return setupDatabase, setupError
}

// -
// Internal
// -

func init() {
	bot.AddEventHandler(onReady)
	bot.AddEventHandler(onDisconnect)
}

func onReady(s *discordgo.Session, _ *discordgo.Ready) {
	db, err := Setup();
	if err != nil {
		utils.Log.Fatalln("Failed to connect to the database:", err)
		return
	}
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := db.Client().Ping(ctx, nil); err != nil {
		utils.Log.Fatalln("Failed to check the database connection:", err)
		return
	}

	utils.Log.Println("Successfully connected to the database!")
}

func onDisconnect(s *discordgo.Session, _ *discordgo.Disconnect) {
	if setupDatabase != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
		defer cancel()

		setupDatabase.Client().Disconnect(ctx)
	}
}
